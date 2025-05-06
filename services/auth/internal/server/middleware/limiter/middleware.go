package limiter

import (
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisstore "github.com/ulule/limiter/v3/drivers/store/redis"
	"time"
)

type Middleware struct {
	cfg   *config.Config
	clnts *clients.Clients
}

func NewMiddleware(cfg *config.Config, clnts *clients.Clients) *Middleware {
	return &Middleware{
		cfg:   cfg,
		clnts: clnts,
	}
}

func (m *Middleware) RateLimiterMiddleware() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  5,
	}

	store, err := redisstore.NewStoreWithOptions(m.clnts.RedisClnt.Redis, limiter.StoreOptions{
		Prefix: "limiter",
	})
	if err != nil {
		panic(err)
	}
	rateLimiter := limiter.New(store, rate)

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		limCtx, err := rateLimiter.Get(ctx, ip)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Rate limiter error"})
			return
		}

		ctx.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limCtx.Limit))
		ctx.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", limCtx.Remaining))
		ctx.Header("X-RateLimit-Reset", fmt.Sprintf("%d", limCtx.Reset))

		if limCtx.Reached {
			ctx.AbortWithStatusJSON(429, gin.H{"error": "Too many requests", "ip": ip})
			return
		}
		ctx.Next()
	}
}
