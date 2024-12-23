package sender

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/sync/errgroup"
)

type Sender interface {
	Connect(ctx context.Context, ssid string, password string) error
}

type sender struct {
	urls []string
}

func New(urls string) Sender {
	u := make([]string, 0, 10)

	for _, url := range strings.Split(urls, " ") {
		if url != "" {
			u = append(u, "http://"+url+":19050")
		}
	}

	return &sender{
		urls: u,
	}
}

func (s *sender) Connect(ctx context.Context, ssid string, password string) error {
	errg := errgroup.Group{}

	for _, url := range s.urls {
		errg.Go(func() error {
			return s.tryToConnect(ctx, url, ssid, password)
		})
	}

	return errg.Wait()
}

type request struct {
	Ssid     string `json:"SSID"`
	Password string `json:"Password"`
}

func (s *sender) tryToConnect(ctx context.Context, url string, ssid string, password string) error {
	req := request{
		Ssid:     ssid,
		Password: password,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("sender.tryToConnect: %w", err)
	}
	fmt.Println(string(data))

	_, _ = http.Post(
		fmt.Sprintf("%s/connect", url),
		"application/json",
		bytes.NewReader(data),
	)

	return nil
}
