package sensors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Temperature struct {
	BaseSensor
	apiUrl string
}

func NewTemperature(id string, apiUrl string) *Temperature {
	return &Temperature{
		BaseSensor: BaseSensor{
			id:         id,
			type_:      TemperatureType,
			lastUpdate: time.Now(),
		},
		apiUrl: apiUrl,
	}
}

func (te *Temperature) Collect() (float64, string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/data", te.apiUrl))
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	var data struct {
		Temperature float64 `json:"temperature"`
		Pressure float64 `json:"pressure"`
		Name string `json:"deviceName"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return -1, "", err
	}
	fmt.Println(data)

	te.lastUpdate = time.Now()

	return data.Temperature, te.type_, nil
}
