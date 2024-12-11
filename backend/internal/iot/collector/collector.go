package collector

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/4aykovski/iot-hub/backend/internal/iot/collector/sensors"
	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
)

type DataRepository interface {
	SaveData(ctx context.Context, data model.Data) error
}

type DeviceRepository interface {
	GetDevices(ctx context.Context) ([]model.Device, error)
}

type Collector struct {
	sensors          []sensors.Sensor
	dataRepository   DataRepository
	deviceRepository DeviceRepository
	interval         time.Duration
}

func New(
	sensors []sensors.Sensor,
	dataRepository DataRepository,
	deviceRepository DeviceRepository,
	interval time.Duration,
) *Collector {
	return &Collector{
		sensors:          sensors,
		dataRepository:   dataRepository,
		deviceRepository: deviceRepository,
		interval:         interval,
	}
}

func (c *Collector) Start(ctx context.Context) error {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			c.collectData()
		}
	}
}

func (c *Collector) collectData() {
	var wg sync.WaitGroup
	for _, s := range c.sensors {
		wg.Add(1)
		go func(s sensors.Sensor) {
			defer wg.Done()
			value, type_, err := s.Collect()
			if err != nil {
				log.Printf("Error collecting data from sensor %s: %v", s.ID(), err)
				return
			}
			fmt.Println(value, type_)

			err = c.dataRepository.SaveData(context.Background(), model.Data{
				Timestamp: time.Now(),
				Value:     value,
				DeviceID:  type_,
			})
			if err != nil {
				log.Printf("Error saving data for sensor %s: %v", s.ID(), err)
			}
		}(s)
	}
	wg.Wait()
}
