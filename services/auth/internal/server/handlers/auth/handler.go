package auth

import (
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/DanHerasymenko/GoDelivery/shared/validator"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

type SignUpReqBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12,max=32"`
}

// @Summary SingUp
// @Description SingUp
// @Tags Auth
// @Param body body SignUpReqBody true "SingUp request body"
// @Success 200 {string} string "SingUp success"
// @Failure 409 {string} string "user already exists"
// @Router /api/auth/singup [post]
func (h *Handler) SingUp(ctx *gin.Context) {

	validator.

	ctx.JSON(200, "SingUp success")
}