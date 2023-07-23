package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Host          string `envconfig:"APPLICATION_HOST" default:"0.0.0.0"`
	Port          string `envconfig:"APPLICATION_PORT" default:"8080"`
	ImmuDbBaseUrl string `envconfig:"IMMUDB_BASE_URL" required:"true"`
	ImmuDbAPIKey  string `envconfig:"IMMUDB_API_KEY"`
}

func ReadConfig() AppConfig {
	var config AppConfig
	err := envconfig.Process("", &config)
	if err != nil {
		panic(fmt.Errorf("couldn't read config %w", err))
	}

	return config
}
