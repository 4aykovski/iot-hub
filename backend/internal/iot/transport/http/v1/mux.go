package v1

import (
	"github.com/4aykovski/iot-hub/backend/internal/iot/transport/http/v1/handlers"
	"github.com/go-chi/chi/v5"
	chiCors "github.com/go-chi/cors"
)

func New(
	deviceHandler *handlers.DeviceHandler,
	dataHandler *handlers.Data,
) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(chiCors.Handler(chiCors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	}))

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/devices", deviceHandler.GetDevices())

		r.Route("/devices/{id}", func(r chi.Router) {
			r.Get("/", deviceHandler.GetDevice())
			r.Put("/", deviceHandler.UpdateDevice())
			r.Get("/data", dataHandler.GetDeviceData())
		})
	})

	return mux
}
