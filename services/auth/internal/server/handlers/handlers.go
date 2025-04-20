package handlers

import (
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/handlers/healthz"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Healthz *healthz.Handler
}

func NewHandlers(cfg *config.Config) *Handlers {
	return &Handlers{
		Healthz: healthz.NewHandler(cfg),
	}
}

func (h *Handlers) RegisterRoutes(router *gin.Engine) {
	router.GET("/healthz", h.Healthz.HealthzCheck)
}
