package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/4aykovski/iot-hub/backend/internal/iot/collector"
	"github.com/4aykovski/iot-hub/backend/internal/iot/collector/sensors"
	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/repo/repoerrs"
	v1 "github.com/4aykovski/iot-hub/backend/internal/iot/transport/http/v1"
	"golang.org/x/sync/errgroup"
)

type App struct {
	provider *Provider

	httpServer *http.Server
}

func NewIotApp(ctx context.Context) *App {
	var a App

	a.initDeps(ctx)

	return &a
}

func (a *App) Start(ctx context.Context) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		slog.Info("iot server started", slog.String("addr", a.httpServer.Addr))
		err := a.startHttp(ctx)
		return err
	})

	eg.Go(func() error {
		slog.Info("iot collector started")
		err := a.startCollector(ctx)
		return err
	})

	return eg.Wait()
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	slog.Info("iot server stopped")

	return nil
}

func (a *App) initDeps(ctx context.Context) {
	inits := []func(context.Context) error{
		a.initProvider,
		a.initCollector,
		a.initHttp,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			panic(err)
		}
	}
}

func (a *App) initProvider(ctx context.Context) error {
	a.provider = NewProvider()

	return nil
}

func (a *App) initCollector(ctx context.Context) error {
	u := strings.ReplaceAll(a.provider.Config().URLs, "'", "")
	u = strings.ReplaceAll(u, " ", "")

	tempSensor := sensors.NewTemperature(
		"temperature-sensor",
		fmt.Sprintf("http://%s:19050", u),
	)
	pressureSensor := sensors.NewPressure(
		"pressure-sensor",
		fmt.Sprintf("http://%s:19050", u),
	)

	_, err := a.provider.DeviceRepository(ctx).GetDevice(ctx, tempSensor.ID())
	if err != nil {
		switch {
		case errors.Is(err, repoerrs.ErrNoDevice):
			a.provider.DeviceRepository(ctx).CreateDevice(ctx, model.Device{
				ID:    tempSensor.ID(),
				Name:  "temperature",
				Limit: -1,
				Type:  tempSensor.Type(),
			})
		default:
			panic(err)
		}
	}

	_, err = a.provider.DeviceRepository(ctx).GetDevice(ctx, pressureSensor.ID())
	if err != nil {
		switch {
		case errors.Is(err, repoerrs.ErrNoDevice):
			a.provider.DeviceRepository(ctx).CreateDevice(ctx, model.Device{
				ID:    pressureSensor.ID(),
				Name:  "pressure",
				Limit: -1,
				Type:  pressureSensor.Type(),
			})
		default:
			panic(err)
		}
	}

	sensors := []sensors.Sensor{
		tempSensor,
		pressureSensor,
	}

	a.provider.SetCollector(collector.New(
		sensors,
		a.provider.DataRepository(ctx),
		a.provider.DeviceRepository(ctx),
		a.provider.Config().Interval,
	))

	return nil
}

func (a *App) initHttp(ctx context.Context) error {
	mux := v1.New(a.provider.DeviceHandler(ctx), a.provider.DataHandler(ctx))

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.provider.Config().Host, a.provider.Config().Port),
		Handler: mux,
	}

	return nil
}

func (a *App) startHttp(ctx context.Context) error {
	if err := a.httpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}

func (a *App) startCollector(ctx context.Context) error {
	return a.provider.Collector(ctx).Start(ctx)
}
