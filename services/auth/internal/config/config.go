package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"time"
)

type Config struct {
	AuthEnv      string `env:"AUTH_APP_ENV"   envDefault:"local"`
	AuthHostPort string `env:"AUTH_HOST_PORT" envDefault:":8080"`

	PostgresContainerHost string `env:"POSTGRES_CONTAINER_HOST"`
	PostgresContainerPort int    `env:"POSTGRES_CONTAINER_PORT"`
	AuthPostgresUser      string `env:"AUTH_POSTGRES_USER"`
	AuthPostgresPassword  string `env:"AUTH_POSTGRES_PASSWORD"`
	AuthPostgresDB        string `env:"AUTH_POSTGRES_DB"`
	RunMigrations         bool   `env:"RUN_MIGRATIONS" envDefault:"false"`

	RedisAddr     string `env:"REDIS_ADDR"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`

	TokenSecret         string        `env:"TOKEN_SECRET"`
	AccessTokenTTLMin   time.Duration `env:"ACCESS_TOKEN_TTL_MIN"`
	RefreshTokenTTLDays time.Duration `env:"REFRESH_TOKEN_TTL_DAYS"`
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from env: %w", err)
	}
	return cfg, nil
}
