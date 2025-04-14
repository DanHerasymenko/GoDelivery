package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Env          string `env:"APP_ENV" envDefault:"local"`
	AuthSrvcPort string `env:"AUTH_SERVICE_PORT" envDefault:":8080"`
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from env: %w", err)
	}
	return cfg, nil
}
