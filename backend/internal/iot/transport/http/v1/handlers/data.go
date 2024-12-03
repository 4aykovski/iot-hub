package handlers

import (
	"context"
	"net/http"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
)

type DataService interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
}

type Data struct {
	dataService DataService
}

func NewData(dataService DataService) *Data {
	return &Data{
		dataService: dataService,
	}
}

type GetDeviceData struct {
	Data []model.Data `json:"data"`
}

func (h *Data) GetDeviceData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
