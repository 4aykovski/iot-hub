package service

import (
	"context"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
)

type DataRepository interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
}

type Data struct {
	dataRepo DataRepository
}

func NewData(dataRepo DataRepository) *Data {
	return &Data{dataRepo: dataRepo}
}
