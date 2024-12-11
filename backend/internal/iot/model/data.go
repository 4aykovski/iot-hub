package model

import "time"

type Data struct {
	ID        int
	Timestamp time.Time
	Value     float64
	DeviceID  string
}
