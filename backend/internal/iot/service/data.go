package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/repo/repoerrs"
)

type DataRepository interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
	GetDeviceDataForPeriod(
		ctx context.Context,
		id string,
		start, end time.Time,
	) ([]model.Data, error)
}

type Data struct {
	dataRepo DataRepository
}

func (da *Data) GetDeviceData(ctx context.Context, id string, interval int) ([]model.Data, error) {
	if interval <= 0 {
		return da.dataRepo.GetDeviceData(ctx, id)
	}

	dateTo := time.Now()
	dateFrom := dateTo.Add(time.Duration(interval) * -1 * time.Second)

	fmt.Println(dateFrom, dateTo)

	return da.dataRepo.GetDeviceDataForPeriod(ctx, id, dateFrom, dateTo)
}

type GetDataForPeriodDTO struct {
	ID   string
	From time.Time
	To   time.Time
}

func (da *Data) GetDataFromPeriod(
	ctx context.Context,
	dto GetDataForPeriodDTO,
) ([]model.Data, error) {
	data, err := da.dataRepo.GetDeviceDataForPeriod(ctx, dto.ID, dto.From, dto.To)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNoData) {
			return []model.Data{}, ErrNoData
		}

		return nil, err
	}
	return data, nil
}

func (da *Data) SaveData(ctx context.Context, data model.Data) error {
	panic("not implemented") // TODO: Implement
}

func NewData(dataRepo DataRepository) *Data {
	return &Data{dataRepo: dataRepo}
}
