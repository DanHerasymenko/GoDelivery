package clients

import (
	"context"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients/postgresClient"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients/redisClient"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
)

type Clients struct {
	PostgresClnt *postgresClient.Client
	RedisClnt    *redisClient.Client
}

func NewClients(ctx context.Context, cfg *config.Config) (*Clients, error) {

	postgresClient, err := postgresClient.NewPostgresClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	redisClint, err := redisClient.NewRedisClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	clients := &Clients{
		PostgresClnt: postgresClient,
		RedisClnt:    redisClint,
	}

	return clients, nil

}
