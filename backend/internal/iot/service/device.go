package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/repo/repoerrs"
)

type DeviceRepository interface {
	GetDevices(ctx context.Context) ([]model.Device, error)
	GetDevice(ctx context.Context, id string) (model.Device, error)
	UpdateDevice(ctx context.Context, device model.Device) error
}

type Device struct {
	deviceRepo DeviceRepository
}

func (de *Device) GetDevices(ctx context.Context) ([]model.Device, error) {
	devices, err := de.deviceRepo.GetDevices(ctx)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNoDevice) {
			return []model.Device{}, nil
		}

		return nil, fmt.Errorf("device.GetDevices: %w", err)
	}

	return devices, nil
}

func (de *Device) GetDevice(ctx context.Context, id string) (model.Device, error) {
	device, err := de.deviceRepo.GetDevice(ctx, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNoDevice) {
			return model.Device{}, ErrNoDevice
		}

		return model.Device{}, fmt.Errorf("device.GetDevice: %w", err)
	}

	return device, nil
}

func (de *Device) UpdateDevice(ctx context.Context, device model.Device) error {
	err := de.deviceRepo.UpdateDevice(ctx, device)
	if err != nil {
		return fmt.Errorf("device.UpdateDevice: %w", err)
	}

	return nil
}

func NewDevice(deviceRepo DeviceRepository) *Device {
	return &Device{deviceRepo: deviceRepo}
}
