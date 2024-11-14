package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Http
}

type Http struct {
	Host string `env:"CONNECTOR_HTTP_HOST" env-default:"0.0.0.0"`
	Port string `env:"CONNECTOR_HTTP_PORT" env-default:"18080"`
}

func Load() *Config {
	var cfg Config

	err := godotenv.Load("./configs/.env.connector")
	if err != nil {
		panic(err.Error())
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err.Error())
	}

	return &cfg
}
