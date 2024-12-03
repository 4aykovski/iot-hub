package handlers

import (
	"context"
	"net/http"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
)

type DeviceService interface {
	GetDevices(ctx context.Context) ([]model.Device, error)
	GetDevice(ctx context.Context, id string) (model.Device, error)
	UpdateDevice(ctx context.Context, device model.Device) error
}

type DeviceHandler struct {
	deviceService DeviceService
}

func NewDevice(deviceService DeviceService) *DeviceHandler {
	return &DeviceHandler{
		deviceService: deviceService,
	}
}

type GetDevicesResponse struct {
	Devices []model.Device `json:"devices"`
}

func (h *DeviceHandler) GetDevices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

type GetDeviceResponse struct {
	Device model.Device `json:"device"`
}

func (h *DeviceHandler) GetDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *DeviceHandler) UpdateDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
