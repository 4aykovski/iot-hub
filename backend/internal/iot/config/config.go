package config

import (
	"github.com/4aykovski/iot-hub/backend/pkg/database/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Http

	Postgres postgres.Config
}

type Http struct {
	Host string `env:"IOT_HTTP_HOST" env-default:"0.0.0.0"`
	Port string `env:"IOT_HTTP_PORT" env-default:"18081"`
}

func Load() *Config {
	var cfg Config

	err := godotenv.Load("./configs/.env.iot")
	if err != nil {
		panic(err.Error())
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err.Error())
	}

	return &cfg
}
