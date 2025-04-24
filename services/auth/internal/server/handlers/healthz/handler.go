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
	Env    string `json:"env" validate:"required"`
}

// HealthzCheck godoc
// @Summary Healthz check
// @Description Healthz check
// @Tags Healthz
// @Accept json
// @Produce json
// @Success      200  {object}  HealthzRespBody
// @Failure 500 {string} string "internal server error"
// @Router /api/healthz [get]
func (h *Handler) HealthzCheck(ctx *gin.Context) {
	ctx.JSON(200, HealthzRespBody{
		Status: "ok",
		Env:    h.cfg.Env,
	})
}
