package model

import "time"

type Data struct {
	ID          string
	Timestamp   time.Time
	Temperature float64
	Humidity    float64
	DeviceID    string
}
