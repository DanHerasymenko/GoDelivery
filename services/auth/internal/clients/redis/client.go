package redis

import (
	"context"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Redis *redis.Client
}

func NewRedisClient(ctx context.Context, cfg *config.Config) (*Client, error) {

	redisURL := fmt.Sprintf("redis://:%s@%s/%d", cfg.RedisPassword, cfg.RedisAddr, cfg.RedisDB)

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	redisClient := redis.NewClient(opt)

	// test connection
	if err := pingCheck(ctx, redisClient); err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %w", err)

	}

	return &Client{Redis: redisClient}, nil
}

func pingCheck(ctx context.Context, redisClient *redis.Client) error {

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return err
	} else if err == nil {
		logger.Info(ctx, "Redis ping successful")
		return nil
	}
	return nil
}
