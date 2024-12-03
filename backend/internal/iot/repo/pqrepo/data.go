package pqrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/repo/repoerrs"
	"github.com/4aykovski/iot-hub/backend/pkg/database/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Data struct {
	DB postgres.DB
}

func NewData(db postgres.DB) *Data {
	return &Data{DB: db}
}

func (d *Data) GetDeviceData(ctx context.Context, id string) ([]model.Data, error) {
	conn := d.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Select("*").From("data").Where("device_id = ?", id).ToSql()
	if err != nil {
		return []model.Data{}, fmt.Errorf("data.GetDeviceData: %w", err)
	}

	var data []model.Data
	rows, err := conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.Data{}, repoerrs.ErrNoData
		}

		return []model.Data{}, fmt.Errorf("data.GetDeviceData: %w", err)
	}

	for rows.Next() {
		var datum model.Data
		err = rows.Scan(
			&datum.ID,
			&datum.Timestamp,
			&datum.Temperature,
			&datum.Humidity,
			&datum.DeviceID,
		)
		if err != nil {
			return []model.Data{}, fmt.Errorf("data.GetDeviceData: %w", err)
		}
		data = append(data, datum)
	}

	if rows.Err(); err != nil {
		return []model.Data{}, fmt.Errorf("data.GetDeviceData: %w", err)
	}

	return data, nil
}
