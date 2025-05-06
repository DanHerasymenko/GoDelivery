package auth

import (
	"errors"
	"fmt"
	_ "fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/constants"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/services"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
	"github.com/DanHerasymenko/GoDelivery/shared/utils/response"
	"github.com/DanHerasymenko/GoDelivery/shared/utils/validator"
	"github.com/gin-gonic/gin"
	"log/slog"
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

	passHash, err := h.services.Auth.GeneratePasswordHash(reqBody.Password)
	if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	_, err = h.services.Auth.CreateUser(ctx, reqBody.Email, passHash)
	if errors.Is(err, constants.ErrUserAlreadyExists) {
		response.AbortWithErrorJSON(ctx, http.StatusConflict, err, err.Error())
		return
	} else if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SingUp success",
	})
}

type SignInReqBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12,max=32"`
}

type SignInResp200Body struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}

// @Summary SignIn
// @Description SignIn
// @Tags Auth
// @Param body body SignInReqBody true "SignIn request body"
// @Success 200 {object} SignInResp200Body "SingIn success"
// @Failure 401 {string} string "invalid credentials"
// @Router /api/auth/signin [post]
func (h *Handler) SignIn(ctx *gin.Context) {

	reqBody := SignInReqBody{}

	if err := validator.ParseReqBody(ctx, &reqBody); err != nil {
		response.AbortWithErrorJSON(ctx, http.StatusBadRequest, err, "Email or password is invalid")
		return
	}

	user, err := h.services.Auth.GetUserByEmail(ctx, reqBody.Email)
	if errors.Is(err, constants.ErrUserNotFound) {
		response.AbortWithErrorJSON(ctx, http.StatusUnauthorized, err, "Invalid credentials")
		return
	}

	ok, err := h.services.Auth.ComparePasswordHashAndPassword(user.PasswordHash, reqBody.Password)
	if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}
	if !ok {
		response.AbortWithErrorJSON(ctx, http.StatusUnauthorized, fmt.Errorf("compare password = false"), "Invalid credentials")
		return
	}

	logger.GinSetLoggerAttr(ctx, slog.String("userID", user.Id))

	accessToken, err := h.services.Auth.CreateAccessAuthToken(ctx.Request.Context(), user.Id)
	if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := h.services.Auth.CreateRefreshAuthToken(ctx.Request.Context(), user.Id)
	if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, SignInResp200Body{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

type RefreshReqBody struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshResp200Body struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}

// @Summary Refresh
// @Description Refresh
// @Tags Auth
// @Param body body RefreshReqBody true "Refresh request body"
// @Success 200 {object} RefreshResp200Body "Refresh success"
// @Failure 401 {string} string "invalid credentials"
// @Failure 500 {string} string "internal server error"
// @Router /api/auth/refresh [post]
func (h *Handler) Refresh(ctx *gin.Context) {

	reqBody := &RefreshReqBody{}

	if err := validator.ParseReqBody(ctx, reqBody); err != nil {
		response.AbortWithError(ctx, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := h.services.Auth.VerifyJWTToken(reqBody.RefreshToken)
	if err != nil {
		response.AbortWithError(ctx, http.StatusUnauthorized, err)
		return
	}

	logger.GinSetLoggerAttr(ctx, slog.String("userID", userIdFromToken))

	if exists, err := h.services.Auth.IfRefreshTokenExistsInRedis(ctx.Request.Context(), reqBody.RefreshToken, userIdFromToken); err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	} else if !exists {
		response.AbortWithError(ctx, http.StatusUnauthorized, fmt.Errorf("refresh token not found"))
		return
	}

	accessToken, err := h.services.Auth.CreateAccessAuthToken(ctx.Request.Context(), userIdFromToken)
	if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := h.services.Auth.CreateRefreshAuthToken(ctx.Request.Context(), userIdFromToken)
	if err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, RefreshResp200Body{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}

// @Summary Logout
// @Description Logout
// @Tags Auth
// @Success 200 {string} string "Logout success"
// @Failure 401 {string} string "access token not found"
// @Failure 500 {string} string "internal server error"
// @Router /api/auth/logout [delete]
func (h *Handler) Logout(ctx *gin.Context) {

	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" {
		response.AbortWithError(ctx, http.StatusUnauthorized, fmt.Errorf("access token not found"))
		return
	}

	fmt.Println(accessToken)

	userIdFromToken, err := h.services.Auth.VerifyJWTToken(accessToken)
	if err != nil {
		response.AbortWithError(ctx, http.StatusUnauthorized, err)
		return
	}

	logger.GinSetLoggerAttr(ctx, slog.String("userID", userIdFromToken))

	if err := h.services.Auth.DeleteTokenFromRedis(ctx.Request.Context(), userIdFromToken); err != nil {
		response.AbortWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}
