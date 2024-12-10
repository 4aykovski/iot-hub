package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

func main() {
	http.HandleFunc("/data", dataHandler())
	http.HandleFunc("/connect", connectHandler())

	fmt.Println("Listening on port 19050")

	err := http.ListenAndServe("0.0.0.0:19050", nil)
	if err != nil {
		panic(err)
	}
}

type DataResponse struct {
	Name        string  `json:"sensor_name"`
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
}

func dataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request")

		var resp DataResponse

		temp, err := rand.Int(rand.Reader, big.NewInt(50))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		humidity, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Name = "temp-sensor"
		resp.Temperature = float64(temp.Int64())
		resp.Pressure = float64(humidity.Int64())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func connectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
