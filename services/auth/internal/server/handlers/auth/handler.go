package auth

import (
	_ "fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/shared/utils/response"
	"github.com/DanHerasymenko/GoDelivery/shared/utils/validator"
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
// @Router /api/auth/signup [post]
func (h *Handler) SingUp(ctx *gin.Context) {

	reqBody := SignUpReqBody{}

	if err := validator.ParseReqBody(ctx, &reqBody); err != nil {
		response.AbortWithErrorJSON(ctx, http.StatusBadRequest, err, "Email or password is invalid")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SingUp success",
	})
}
