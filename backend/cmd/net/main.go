package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func scanNetwork(subnet string, port string, path string) {
	var wg sync.WaitGroup
	results := make(chan string, 255)

	// Параллельное сканирование
	for i := 1; i <= 254; i++ {
		wg.Add(1)
		go func(host int) {
			defer wg.Done()

			ip := fmt.Sprintf("%s.%d", subnet, host)
			url := fmt.Sprintf("http://%s:%s%s", ip, port, path)

			client := http.Client{
				Timeout: 2 * time.Second, // Короткий таймаут
			}

			fmt.Println(url)
			resp, err := client.Get(url)
			fmt.Println(resp)
			if err == nil {
				defer resp.Body.Close()
				if resp.StatusCode == 200 {
					results <- ip
				}
			}
			fmt.Println(err)
		}(i)
	}

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Сбор результатов
	for result := range results {
		fmt.Printf("Found device at: %s\n", result)
	}
}

type SomeBody struct {
	SSID     string `json:"SSID"`
	Password string `json:"Password"`
}

func main() {
	// Примеры использования
	// scanNetwork("192.168.182", "19050", "/connect") // Типичная домашняя сеть

	// resp, err := http.Post("http://192.168.182.187:19050/connect", "application/json", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println(resp)

	body := SomeBody{
		SSID:     "iPhone",
		Password: "Lionart724",
	}

	str, err := json.Marshal(&body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(str))
	res, err := http.Post(
		"http://192.168.182.187:19050/connect",
		"application/json",
		bytes.NewReader(str),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
