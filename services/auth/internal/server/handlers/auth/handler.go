package auth

import (
	_ "fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/services"
	"github.com/DanHerasymenko/GoDelivery/shared/utils/response"
	"github.com/DanHerasymenko/GoDelivery/shared/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cfg      *config.Config
	services *services.Services
}

func NewHandler(cfg *config.Config, services *services.Services) *Handler {
	return &Handler{
		cfg:      cfg,
		services: services,
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

	user, err := h.services.Auth.CreateUser(ctx, reqBody.Email, reqBody.Password)
	if err != nil {
		response.AbortWithErrorJSON(ctx, http.StatusConflict, err, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SingUp success",
		"user":    user,
	})
	//ctx.JSON(http.StatusOK, gin.H{
	//	"message": "SingUp success",
	//})
}
