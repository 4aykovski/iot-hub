package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	sensor = "temp-sensor"
)

type Response struct {
	Name string `json:"sensor_name"`
}

func scanNetwork(results chan string, subnet string, port string, path string) {
	var wg sync.WaitGroup

	// Параллельное сканирование
	for i := 0; i <= 254; i++ {
		wg.Add(1)
		go func(host int) {
			defer wg.Done()

			ip := fmt.Sprintf("%s.%d", subnet, host)
			url := fmt.Sprintf("http://%s:%s%s", ip, port, path)

			client := http.Client{
				Timeout: 3 * time.Second, // Короткий таймаут
			}

			resp, err := client.Get(url)
			if err == nil {
				defer resp.Body.Close()

				res := Response{}
				err := json.NewDecoder(resp.Body).Decode(&res)
				if err != nil {
					return
				}

				if resp.StatusCode == 200 && res.Name == sensor {
					results <- ip
				}
			}
		}(i)
	}

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()
}

func main() {
	subnet := flag.String("subnet", "127.0.0", "Subnet to scan")
	port := flag.String("port", "8080", "Port to scan")
	path := flag.String("path", "/data", "Path to scan")
	output := flag.String(
		"output",
		"/home/root/apps/iot-hub/backend/configs/.env.device",
		"Output file",
	)
	flag.Parse()

	fmt.Println("network scan started on", *subnet, *port, *path)

	results := make(chan string, 255)
	scanNetwork(results, *subnet, *port, *path) // Типичная домашняя сеть

	foundNets := "DEVICES_NETWORKS='"

	for result := range results {
		fmt.Printf("Found device at: %s\n", result)
		foundNets += result + " "
	}

	foundNets = foundNets + "'\n"

	outputFile, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(foundNets)
	if err != nil {
		panic(err)
	}

	fmt.Println("network scan done")
}
