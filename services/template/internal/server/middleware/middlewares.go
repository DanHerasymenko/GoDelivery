package middleware

import (
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/middleware/logging"
	_ "github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/middleware/logging"
)

type Middlewares struct {
	Log *logging.Middleware
}

func NewMiddlewares(cfg *config.Config, clnts *clients.Clients) *Middlewares {
	return &Middlewares{
		Log: logging.NewMiddleware(cfg),
	}
}
