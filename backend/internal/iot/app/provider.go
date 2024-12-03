package app

import (
	"context"
	"time"

	"github.com/4aykovski/iot-hub/backend/internal/iot/config"
	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/repo/pqrepo"
	"github.com/4aykovski/iot-hub/backend/internal/iot/service"
	"github.com/4aykovski/iot-hub/backend/internal/iot/transport/http/v1/handlers"
	"github.com/4aykovski/iot-hub/backend/pkg/database/postgres"
)

type DataRepository interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
	GetDeviceDataForPeriod(
		ctx context.Context,
		id string,
		start, end time.Time,
	) ([]model.Data, error)
	SaveData(ctx context.Context, data model.Data) error
}

type DeviceRepository interface {
	GetDevices(ctx context.Context) ([]model.Device, error)
	GetDevice(ctx context.Context, id string) (model.Device, error)
	UpdateDevice(ctx context.Context, device model.Device) error
	CreateDevice(ctx context.Context, device model.Device) error
}

type DataService interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
	GetDataFromPeriod(ctx context.Context, dto service.GetDataForPeriodDTO) ([]model.Data, error)
	SaveData(ctx context.Context, data model.Data) error
}

type DeviceService interface {
	GetDevices(ctx context.Context) ([]model.Device, error)
	GetDevice(ctx context.Context, id string) (model.Device, error)
	UpdateDevice(ctx context.Context, device model.Device) error
}

type Collector interface {
	Start(ctx context.Context) error
}

type Provider struct {
	config    *config.Config
	db        *postgres.DB
	collector Collector

	// repositories
	dataRepository   DataRepository
	deviceRepository DeviceRepository

	// services
	dataService   DataService
	deviceService DeviceService

	// handlers
	dataHandler   *handlers.Data
	deviceHandler *handlers.DeviceHandler
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Config() *config.Config {
	if p.config == nil {
		p.config = config.Load()
	}

	return p.config
}

func (p *Provider) DB(ctx context.Context) *postgres.DB {
	if p.db == nil {
		db, err := postgres.New(ctx, p.Config().Postgres)
		if err != nil {
			panic(err)
		}

		p.db = db
	}

	return p.db
}

func (p *Provider) DataRepository(ctx context.Context) DataRepository {
	if p.dataRepository == nil {
		p.dataRepository = pqrepo.NewData(*p.DB(ctx))
	}

	return p.dataRepository
}

func (p *Provider) DeviceRepository(ctx context.Context) DeviceRepository {
	if p.deviceRepository == nil {
		p.deviceRepository = pqrepo.NewDevice(*p.DB(ctx))
	}

	return p.deviceRepository
}

func (p *Provider) DataService(ctx context.Context) DataService {
	if p.dataService == nil {
		p.dataService = service.NewData(p.DataRepository(ctx))
	}

	return p.dataService
}

func (p *Provider) DeviceService(ctx context.Context) DeviceService {
	if p.deviceService == nil {
		p.deviceService = service.NewDevice(p.DeviceRepository(ctx))
	}

	return p.deviceService
}

func (p *Provider) DataHandler(ctx context.Context) *handlers.Data {
	if p.dataHandler == nil {
		p.dataHandler = handlers.NewData(p.DataService(ctx))
	}

	return p.dataHandler
}

func (p *Provider) DeviceHandler(ctx context.Context) *handlers.DeviceHandler {
	if p.deviceHandler == nil {
		p.deviceHandler = handlers.NewDevice(p.DeviceService(ctx))
	}

	return p.deviceHandler
}

func (p *Provider) SetCollector(collector Collector) {
	p.collector = collector
}

func (p *Provider) Collector(ctx context.Context) Collector {
	if p.collector == nil {
		return nil
	}

	return p.collector
}
