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

type Device struct {
	DB postgres.DB
}

func (de *Device) CreateDevice(ctx context.Context, device model.Device) error {
	conn := de.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Insert("devices").
		Columns("id", "name", "\"limit\"", "type").
		Values(device.ID, device.Name, device.Limit, device.Type).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("device.CreateDevice: %w", err)
	}

	_, err = conn.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("device.CreateDevice: %w", err)
	}

	return nil
}

func NewDevice(db postgres.DB) *Device {
	return &Device{DB: db}
}

func (d *Device) GetDevice(ctx context.Context, id string) (model.Device, error) {
	conn := d.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Select("*").From("devices").Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.Device{}, fmt.Errorf("device.GetDevice: %w", err)
	}

	var device model.Device
	err = conn.QueryRow(ctx, stmt, args...).
		Scan(&device.ID, &device.Name, &device.Limit, &device.Type)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Device{}, repoerrs.ErrNoDevice
		}

		return model.Device{}, fmt.Errorf("device.GetDevice: %w", err)
	}

	return device, nil
}

func (d *Device) GetDevices(ctx context.Context) ([]model.Device, error) {
	conn := d.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Select("*").From("devices").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return []model.Device{}, fmt.Errorf("device.GetDevices: %w", err)
	}

	var devices []model.Device
	rows, err := conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.Device{}, repoerrs.ErrNoDevice
		}

		return []model.Device{}, fmt.Errorf("device.GetDevices: %w", err)
	}

	for rows.Next() {
		var device model.Device
		err = rows.Scan(&device.ID, &device.Name, &device.Limit, &device.Type)
		if err != nil {
			return []model.Device{}, fmt.Errorf("device.GetDevices: %w", err)
		}
		devices = append(devices, device)
	}

	if rows.Err(); err != nil {
		return []model.Device{}, fmt.Errorf("device.GetDevices: %w", err)
	}

	return devices, nil
}

func (d *Device) UpdateDevice(ctx context.Context, device model.Device) error {
	conn := d.DB.GetConnection(ctx)

	stmt, args, err := squirrel.Update("devices").
		Set("name", device.Name).
		Set("\"limit\"", device.Limit).
		Set("type", device.Type).
		Where("id = ?", device.ID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("device.UpdateDevice: %w", err)
	}

	_, err = conn.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("device.UpdateDevice: %w", err)
	}

	return nil
}
