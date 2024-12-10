package controllers

import (
	"EduConnect/pkg/config"
	"EduConnect/pkg/jwt"
	"EduConnect/pkg/logger"
	"EduConnect/pkg/requests"
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	log logger.Logger
	cfg *config.Config
	// accountRepo  repository.AccountRepository
	validate *validator.Validate
	// nonceManager nonce.NonceManager
	jwtManager jwt.JwtManager
}

func NewAuthController(log logger.Logger, cfg *config.Config, validator *validator.Validate, jwtManager jwt.JwtManager) *AuthController {
	return &AuthController{log: log, cfg: cfg, validate: validator, jwtManager: jwtManager}
}

// RefreshTokens godoc
// @Summary Обновление Access и Refresh токенов
// @Description Обновление access и refresh токенов с использованием валидного refresh токена
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   refreshTokens body requests.RefreshTokensRequest true "RefreshTokens Request"
// @Success 200 {object} response.RefreshTokensResponse "Новые access и refresh токены"
// @Failure 400 {object} response.ErrorResponse "Неверный refresh токен или ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/auth/refresh-tokens [post]
func (a *AuthController) RefreshTokens(ctx echo.Context) error {
	a.log.Infof("(AuthController.RefreshTokens)")
	var req requests.RefreshTokensRequest
	if err := a.decodeRequest(ctx, &req); err != nil {
		a.log.Debugf("Failed to validate request RefreshTokens: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	accessToken, refreshToken, err := a.jwtManager.RefreshToken(context.Background(), req.RefreshToken)
	if err != nil {
		a.log.Debugf("(RefreshTokens) error: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshTokens godoc
// @Summary Обновление Access и Refresh токенов
// @Description Обновление access и refresh токенов с использованием валидного refresh токена
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   refreshTokens body requests.SignOutRequest true "RefreshTokens Request"
// @Success 200 {object} response.SignOutResponse "Пустой ответ при успешном выходе"
// @Failure 400 {object} response.ErrorResponse "При неверных токенах доступа"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/auth/refresh-tokens [post]
func (a *AuthController) SignOut(ctx echo.Context) error {
	var req requests.SignOutRequest
	if err := a.decodeRequest(ctx, &req); err != nil {
		a.log.Debugf("Failed to validate SignOutRequest: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)
	token := ctx.Get("token").(string)

	accountId := accountClaims["sub"].(string)
	deviceId := accountClaims["device_id"].(string)

	err := a.jwtManager.RevokeTokens(context.Background(), accountId, deviceId, token)
	if err != nil {
		a.log.Debugf("(SignOut) error by revoking tokens: %v")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error by revoking tokens"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{})
}

// func (a *AuthController) VerifyToken(ctx echo.Context) error {
// 	a.log.Debugf("VerifyToken")
// 	var req requests.VerifyTokenRequest
// 	if err := a.decodeRequest(ctx, &req); err != nil {
// 		a.log.Debugf("Failed to validate VerifyToken: %v", err)
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
// 	}

// 	mapClaims, err := a.jwtManager.VerifyToken(context.Background(), req.AccessToken)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Failed to verify token: %v", err)})
// 	}

// 	sub := mapClaims["sub"].(string)

// 	return ctx.JSON(http.StatusOK, map[string]string{
// 		"account_id": sub,
// 	})
// }
