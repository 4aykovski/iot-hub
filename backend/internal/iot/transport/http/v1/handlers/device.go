package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/service"
	"github.com/4aykovski/iot-hub/backend/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	response.Response
}

func (h *DeviceHandler) GetDevices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		devices, err := h.deviceService.GetDevices(r.Context())
		if err != nil {
			slog.Error(
				"device.GetDevices",
				slog.String("error", err.Error()),
			)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.InternalError())
			return
		}

		render.JSON(w, r, GetDevicesResponse{
			Response: response.OK(),
			Devices:  devices,
		})
	}
}

type GetDeviceResponse struct {
	Device model.Device `json:"device"`
	response.Response
}

func (h *DeviceHandler) GetDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		device, err := h.deviceService.GetDevice(r.Context(), id)
		if err != nil {
			if errors.Is(err, service.ErrNoDevice) {
				slog.Info(
					"device not found",
					slog.String("id", id),
				)

				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, response.NotFoundError())
				return
			}

			slog.Error(
				"device.GetDevice",
				slog.String("error", err.Error()),
				slog.String("id", id),
			)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.InternalError())
			return
		}

		slog.Info(
			"device found",
			slog.String("id", id),
		)

		render.JSON(w, r, GetDeviceResponse{
			Response: response.OK(),
			Device:   device,
		})
	}
}

type UpdateDeviceRequest struct {
	Name  string  `json:"name"`
	Limit float64 `json:"limit"`
}

func (h *DeviceHandler) UpdateDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var req UpdateDeviceRequest
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			slog.Info(
				"invalid request",
				slog.String("id", id),
				slog.String("error", err.Error()),
			)

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.BadRequestError("invalid request"))
			return
		}

		device := model.Device{
			ID:    id,
			Name:  req.Name,
			Limit: req.Limit,
		}

		if err := h.deviceService.UpdateDevice(r.Context(), device); err != nil {
			slog.Error(
				"device.UpdateDevice",
				slog.String("error", err.Error()),
				slog.String("id", id),
			)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.InternalError())
			return
		}

		slog.Info(
			"device updated",
			slog.String("id", id),
		)

		render.JSON(w, r, response.OK())
	}
}
