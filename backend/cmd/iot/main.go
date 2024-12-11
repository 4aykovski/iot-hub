package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/4aykovski/iot-hub/backend/internal/iot/app"
)

func main() {
	ctx := context.Background()
	fmt.Println("starting iot app")

	iot := app.NewIotApp(ctx)
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
