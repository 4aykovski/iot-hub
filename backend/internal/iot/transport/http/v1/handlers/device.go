package handlers

import (
	"net/http"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
)

type DeviceHandler struct {
}

func NewDevice() *DeviceHandler {
	return &DeviceHandler{}
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

type GetDeviceData struct {
	Data []model.Data `json:"data"`
}

func (h *DeviceHandler) GetDeviceData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
