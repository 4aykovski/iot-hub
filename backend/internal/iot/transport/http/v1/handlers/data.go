package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/service"
	"github.com/4aykovski/iot-hub/backend/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type DataService interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
	GetDataFromPeriod(ctx context.Context, dto service.GetDataForPeriodDTO) ([]model.Data, error)
}

type Data struct {
	dataService DataService
}

func NewData(dataService DataService) *Data {
	return &Data{
		dataService: dataService,
	}
}

type GetDeviceDataResponse struct {
	response.Response
	Data []model.Data `json:"data"`
}

func (h *Data) GetDeviceData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

type GetDeviceDataForPeriodResponse struct {
	response.Response
	Data []model.Data `json:"data"`
}

func (h *Data) GetDataForPeriod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		toDateString := r.URL.Query().Get("to")
		fromDateString := r.URL.Query().Get("from")

		if toDateString == "" || fromDateString == "" {
			slog.Info(
				"invalid request",
				slog.String("to", toDateString),
				slog.String("from", fromDateString),
				slog.String("id", id),
			)

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.BadRequestError("invalid request"))
			return
		}

		toDate, err := time.Parse(time.RFC3339, toDateString)
		if err != nil {
			slog.Info(
				"invalid request",
				slog.String("to", toDateString),
				slog.String("from", fromDateString),
				slog.String("id", id),
			)

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.BadRequestError("invalid request"))
			return
		}
		fromDate, err := time.Parse(time.RFC3339, fromDateString)
		if err != nil {
			slog.Info(
				"invalid request",
				slog.String("to", toDateString),
				slog.String("from", fromDateString),
				slog.String("id", id),
			)

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.BadRequestError("invalid request"))
			return
		}

		data, err := h.dataService.GetDataFromPeriod(r.Context(), service.GetDataForPeriodDTO{
			ID:   id,
			From: fromDate,
			To:   toDate,
		})
		if err != nil {
			slog.Error(
				"failed to get data",
				slog.String("error", err.Error()),
				slog.String("id", id),
				slog.Time("from", fromDate),
				slog.Time("to", toDate),
			)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.InternalError())
			return
		}

		slog.Info(
			"got data",
			slog.String("id", id),
			slog.Time("from", fromDate),
			slog.Time("to", toDate),
		)

		render.JSON(w, r, GetDeviceDataForPeriodResponse{
			Data:     data,
			Response: response.OK(),
		})
	}
}
