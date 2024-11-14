package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

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

func (a *App) initHttp(ctx context.Context) error {
	mux := v1.New()

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.provider.Config().Http.Host, a.provider.Config().Http.Port),
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
