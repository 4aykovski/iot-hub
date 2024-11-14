package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	connectorApp "github.com/4aykovski/iot-hub/backend/internal/connector/app"
	iotApp "github.com/4aykovski/iot-hub/backend/internal/iot/app"
)

func main() {
	done := make(chan struct{})
	ctx := context.Background()

	connector := connectorApp.NewConnectorApp(done)
	go func() {
		if err := connector.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			panic(err)
		}
	}()

	<-done
	// test := make(chan os.Signal, 1)
	// signal.Notify(test, syscall.SIGINT, syscall.SIGTERM)
	// <-test
	if err := connector.GracefullStop(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("starting iot app")

	iot := iotApp.NewIotApp(ctx)
	go func() {
		if err := iot.Start(ctx); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := iot.Stop(ctx); err != nil {
		panic(err)
	}

	fmt.Println("done")
}
