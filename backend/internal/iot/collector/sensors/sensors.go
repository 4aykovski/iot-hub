package sensors

import "time"

const (
	TemperatureType = "temperature-sensor"
	HumidityType    = "humidity-sensor"
)

type Sensor interface {
	ID() string
	Collect() (float64, string, error)
	Type() string
	LastUpdate() time.Time
}

type BaseSensor struct {
	id         string
	type_      string
	lastUpdate time.Time
}

func (ba *BaseSensor) Type() string {
	return ba.type_
}

func (ba *BaseSensor) ID() string {
	return ba.id
}

func (ba *BaseSensor) LastUpdate() time.Time {
	return ba.lastUpdate
}
