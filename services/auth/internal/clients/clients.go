package clients

import (
	"context"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients/postgres"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients/redis"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
)

type Clients struct {
	PostgresClnt *postgres.Client
	RedisClnt    *redis.Client
}

func NewClients(ctx context.Context, cfg *config.Config) (*Clients, error) {

	postgresClient, err := postgres.NewPostgresClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	redisClint, err := redis.NewRedisClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	clients := &Clients{
		PostgresClnt: postgresClient,
		RedisClnt:    redisClint,
	}

	return clients, nil

}
