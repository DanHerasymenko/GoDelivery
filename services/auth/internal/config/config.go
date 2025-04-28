package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	AuthEnv              string `env:"AUTH_APP_ENV"   envDefault:"local"`
	AuthHostPort         string `env:"AUTH_HOST_PORT" envDefault:":8080"`
	PostgresHost         string `env:"POSTGRES_HOST"`
	PostgresPort         int    `env:"POSTGRES_PORT"`
	AuthPostgresUser     string `env:"AUTH_POSTGRES_USER"`
	AuthPostgresPassword string `env:"AUTH_POSTGRES_PASSWORD"`
	AuthPostgresDB       string `env:"AUTH_POSTGRES_DB"`
	RunMigrations        bool   `env:"RUN_MIGRATIONS" envDefault:"false"`
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from env: %w", err)
	}
	return cfg, nil
}
