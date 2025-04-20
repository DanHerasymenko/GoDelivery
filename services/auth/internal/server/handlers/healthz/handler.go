package healthz

import (
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

type HealthzRespBody struct {
	Status string `json:"status" validate:"required"`
}

func (h *Handler) HealthzCheck(ctx *gin.Context) {
	ctx.JSON(200, HealthzRespBody{
		Status: "ok",
	})
}
