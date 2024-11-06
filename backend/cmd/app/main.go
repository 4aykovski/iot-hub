package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	connectorApp "github.com/4aykovski/iot-hub/backend/internal/connector/app"
)

func main() {
	done := make(chan struct{})

	connector := connectorApp.New(done)
	go func() {
		if err := connector.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			panic(err)
		}
	}()

	<-done
	if err := connector.GracefullStop(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("done")
}
