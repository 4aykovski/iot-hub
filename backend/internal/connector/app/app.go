package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	v1 "github.com/4aykovski/iot-hub/backend/internal/connector/transport/http/v1"
	"golang.org/x/sync/errgroup"
)

type App struct {
	server *http.Server

	provider *Provider
}

func NewConnectorApp(done chan struct{}) *App {

	a := &App{}

	a.initDeps(done)

	return a
}

func (a *App) Start() error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		slog.Info("connector server started", slog.String("addr", a.server.Addr))
		return a.server.ListenAndServe()
	})

	return eg.Wait()
}

func (a *App) GracefullStop(ctx context.Context) error {
	a.server.Shutdown(ctx)

	slog.Info("connector server stopped")

	return nil
}

func (a *App) initDeps(done chan struct{}) {
	a.initProvider()
	a.initHttp(done)
}

func (a *App) initHttp(done chan struct{}) {
	a.server = &http.Server{
		Addr: fmt.Sprintf(
			"%s:%s",
			a.provider.Config().Http.Host,
			a.provider.Config().Http.Port,
		),
		Handler:      v1.New(done),
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (a *App) initProvider() {
	a.provider = NewProvider()
}
