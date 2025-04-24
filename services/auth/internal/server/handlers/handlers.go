package handlers

import (
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/handlers/healthz"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Healthz *healthz.Handler

	mdlwrs *middleware.Middlewares
}

func NewHandlers(cfg *config.Config, mdlwrs *middleware.Middlewares) *Handlers {
	return &Handlers{
		Healthz: healthz.NewHandler(cfg),

		mdlwrs: mdlwrs,
	}
}

func (h *Handlers) RegisterRoutes(router *gin.Engine) {

	// API group with global log middleware
	api := router.Group("/api")
	api.Use(h.mdlwrs.Log.Handle)

	healtz := api.Group("/healthz")
	healtz.GET("/", h.Healthz.HealthzCheck)
}
