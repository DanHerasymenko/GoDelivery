package redisClient

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCommander interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}
