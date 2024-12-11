package sensors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pressure struct {
	BaseSensor
	apiUrl string
}

func NewPressure(id string, apiUrl string) *Pressure {
	return &Pressure{
		BaseSensor: BaseSensor{
			id:         id,
			type_:      PressureType,
			lastUpdate: time.Now(),
		},
		apiUrl: apiUrl,
	}
}

func (hu *Pressure) Collect() (float64, string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/data", hu.apiUrl))
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	var data struct {
		Pressure float64 `json:"pressure"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return -1, "", err
	}

	hu.lastUpdate = time.Now()

	return data.Pressure, hu.type_, nil
}
