package v1

import (
	"github.com/4aykovski/iot-hub/backend/internal/iot/transport/http/v1/handlers"
	"github.com/go-chi/chi/v5"
)

func New(
	deviceService handlers.DeviceService,
	dataService handlers.DataService,
) *chi.Mux {
	deviceHandler := handlers.NewDevice(deviceService)
	dataHandler := handlers.NewData(dataService)

	mux := chi.NewMux()

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/devices", deviceHandler.GetDevice())

		r.Route("/devices/{id}", func(r chi.Router) {
			r.Get("/", deviceHandler.GetDevice())
			r.Put("/", deviceHandler.UpdateDevice())
		})

		r.Get("/data", dataHandler.GetDeviceData())
	})

	return chi.NewMux()
}
