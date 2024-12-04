package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	sensor = "temp-sensor"
)

type Response struct {
	Name string `json:"sensor_name"`
}

func scanNetwork(
	results chan string,
	subnet string,
	port string,
	path string,
	outWg *sync.WaitGroup,
) {
	defer outWg.Done()
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
	wg.Wait()
}

func main() {
	subnet := flag.String("subnet", "127.0.0.0", "Subnet to scan")
	port := flag.String("port", "8080", "Port to scan")
	path := flag.String("path", "/data", "Path to scan")
	pathToSubnetFile := flag.String(
		"subnetFile",
		"",
		"Path to subnet file",
	)
	output := flag.String(
		"output",
		"/home/root/apps/iot-hub/backend/configs/.env.device",
		"Output file",
	)
	flag.Parse()

	if *pathToSubnetFile != "" {
		file, err := os.Open(*pathToSubnetFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			*subnet = scanner.Text()
		}
	}

	nets := strings.Split(*subnet, " ")
	subnets := make([]string, 0, len(nets))

	for _, net := range nets {
		index := strings.LastIndex(net, ".")

		if index != -1 {
			net = net[:index]
		}

		subnets = append(subnets, net)
	}

	fmt.Println("network scan started on", *port, *path)

	foundNets := "DEVICES_NETWORKS='"
	foundNetsMap := make(map[string]struct{})
netloop:
	for i := 0; i < 3; i++ {
		results := make(chan string, 255)
		wg := sync.WaitGroup{}

		for _, subnet := range subnets {
			wg.Add(1)
			fmt.Println("Scanning subnet", subnet)
			go scanNetwork(results, subnet, *port, *path, &wg)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		for result := range results {
			if _, ok := foundNetsMap[result]; ok {
				continue
			}

			fmt.Printf("Found device at: %s:19050\n", result)
			foundNets += result + " "
			foundNetsMap[result] = struct{}{}
			break netloop
		}

		fmt.Printf("end %d try\n", i)
		fmt.Println("Sleeping for 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	if foundNets == "DEVICES_NETWORKS='" {
		fmt.Println("No devices found")
		return
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
