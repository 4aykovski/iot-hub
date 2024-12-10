package pqrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

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

	stmt, args, err := squirrel.Select("id, device_id, value, timestamp").
		From("data").
		Where("device_id = ?", id).
		OrderBy("timestamp ").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
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
			&datum.DeviceID,
			&datum.Value,
			&datum.Timestamp,
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

func (d *Data) GetDeviceDataForPeriod(
	ctx context.Context,
	id string,
	start time.Time,
	end time.Time,
) ([]model.Data, error) {
	conn := d.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Select("id, device_id, value, timestamp").From("data").
		Where("device_id = ?", id).
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		OrderBy("timestamp ").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return []model.Data{}, fmt.Errorf("data.GetDeviceDataForPeriod: %w", err)
	}

	var data []model.Data
	rows, err := conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.Data{}, repoerrs.ErrNoData
		}

		return []model.Data{}, fmt.Errorf("data.GetDeviceDataForPeriod: %w", err)
	}

	for rows.Next() {
		var datum model.Data
		err = rows.Scan(
			&datum.ID,
			&datum.DeviceID,
			&datum.Value,
			&datum.Timestamp,
		)
		if err != nil {
			return []model.Data{}, fmt.Errorf("data.GetDeviceDataForPeriod: %w", err)
		}
		data = append(data, datum)
	}

	if rows.Err(); err != nil {
		return []model.Data{}, fmt.Errorf("data.GetDeviceDataForPeriod: %w", err)
	}

	return data, nil
}

func (d *Data) SaveData(ctx context.Context, data model.Data) error {
	conn := d.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Insert("data").
		Columns("timestamp", "value", "device_id").
		Values(data.Timestamp, data.Value, data.DeviceID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("data.SaveData: %w", err)
	}

	_, err = conn.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("data.SaveData: %w", err)
	}

	return nil
}
