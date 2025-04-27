package auth

import (
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	ru "github.com/DanHerasymenko/GoDelivery/shared/utils/response"
	su "github.com/DanHerasymenko/GoDelivery/shared/validator"
	"github.com/gin-gonic/gin"
	"net/http"
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
	Email    string `json:"email" validate:"required,email"`
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

	reqBody := SignUpReqBody{}

	// :TODO - check GIN errors and what to return - need to log by slog??
	if err := su.ParseReqBody(ctx, &reqBody); err != nil {

		ru.AbortWithError(ctx, http.StatusBadRequest, fmt.Errorf()
			ctx.AbortWithStatus(http.StatusBadRequest)
		return err
	}

	ctx.JSON(200, reqBody)
}
