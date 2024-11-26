package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	connectorApp "github.com/4aykovski/iot-hub/backend/internal/connector/app"
)

func main() {
	done := make(chan struct{})

	connector := connectorApp.NewConnectorApp(done)
	go func() {
		if err := connector.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			panic(err)
		}
	}()

	test := make(chan os.Signal, 1)
	signal.Notify(test, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-test:
		done <- struct{}{}
	case <-done:
		if err := connector.GracefullStop(context.Background()); err != nil {
			panic(err)
		}
	}
}
