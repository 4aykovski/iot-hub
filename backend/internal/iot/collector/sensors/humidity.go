package sensors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Humidity struct {
	BaseSensor
	apiUrl string
}

func NewHumidity(id string, apiUrl string) *Humidity {
	return &Humidity{
		BaseSensor: BaseSensor{
			id:         id,
			type_:      HumidityType,
			lastUpdate: time.Now(),
		},
		apiUrl: apiUrl,
	}
}

func (hu *Humidity) Collect() (float64, string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/data", hu.apiUrl))
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	var data struct {
		Humidity float64 `json:"pressure"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return -1, "", err
	}

	hu.lastUpdate = time.Now()

	return data.Humidity, hu.type_, nil
}
