package app

import "github.com/4aykovski/iot-hub/backend/internal/iot/config"

type Provider struct {
	config *config.Config
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Config() *config.Config {
	if p.config == nil {
		p.config = config.Load()
	}

	return p.config
}
