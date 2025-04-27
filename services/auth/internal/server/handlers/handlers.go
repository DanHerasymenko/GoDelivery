package handlers

import (
	_ "github.com/DanHerasymenko/GoDelivery/services/auth-service/cmd/server/docs"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	ah "github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/handlers/auth"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/handlers/healthz"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	Healthz *healthz.Handler
	Auth    *ah.Handler

	mdlwrs *middleware.Middlewares
}

func NewHandlers(cfg *config.Config, mdlwrs *middleware.Middlewares) *Handlers {
	return &Handlers{
		Healthz: healthz.NewHandler(cfg),
		Auth:    ah.NewHandler(cfg),

		mdlwrs: mdlwrs,
	}
}

func (h *Handlers) RegisterRoutes(router *gin.Engine) {

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API group with global log middleware
	api := router.Group("/api")
	api.Use(h.mdlwrs.Log.Handle)

	// Healthz group
	healtz := api.Group("/healthz")
	healtz.GET("/", h.Healthz.HealthzCheck)

	// Auth group
	auth := api.Group("/auth")
	auth.POST("/signup", h.Auth.SingUp)
}
