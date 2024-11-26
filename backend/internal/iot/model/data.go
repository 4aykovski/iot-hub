package model

import "time"

type Data struct {
	Timestamp   time.Time
	Temperature float64
	Humidity    float64
}
