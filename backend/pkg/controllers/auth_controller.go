package controllers

import (
	"EduConnect/internal/model"
	"EduConnect/internal/repository"
	"EduConnect/internal/values"
	"EduConnect/pkg/config"
	"EduConnect/pkg/jwt"
	"EduConnect/pkg/logger"
	"EduConnect/pkg/requests"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	log            logger.Logger
	cfg            *config.Config
	validate       *validator.Validate
	jwtManager     jwt.JwtManager
	userRepository repository.UserRepository
}

func NewAuthController(log logger.Logger, cfg *config.Config, userRepository repository.UserRepository, validator *validator.Validate, jwtManager jwt.JwtManager) *AuthController {
	return &AuthController{log: log, cfg: cfg, userRepository: userRepository, validate: validator, jwtManager: jwtManager}
}

// SignUp godoc
// @Summary Регистрация пользователя
// @Description Создаёт нового пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   signUpRequest body requests.SignUpRequest true "Данные для регистрации"
// @Success 200 {object} response.SignUpResponse "Tokens"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/sign-up [post]
func (a *AuthController) SignUp(ctx echo.Context) error {
	a.log.Infof("(AuthController.SignUp)")
	var req requests.SignUpRequest
	if err := a.decodeRequest(ctx, &req); err != nil {
		a.log.Debugf("Failed to validate request SignUp: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	email, err := values.NewEmail(req.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Email error: %v", err)})
	}
	hashedPassword, err := values.NewPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Password error: %v", err)})
	}

	userUuid, _ := uuid.NewV7()
	userId := userUuid.String()
	user := model.NewUser(userId, email, hashedPassword, "Пока так", time.Now(), req.Role)

	deviceUuid, _ := uuid.NewV7()
	deviceId := deviceUuid.String()

	if err := a.userRepository.Create(context.Background(), user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Internal server error by create user: %v", err)})
	}

	accessToken, refreshToken, err := a.jwtManager.GenerateTokens(context.Background(), userId, deviceId, req.Email, req.Role)
	if err != nil {
		a.log.Debugf("(GenerateTokens) Failed to generate tokens: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Internal server error: %v", err)})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// SignIn godoc
// @Summary Вход пользователя
// @Description Авторизация пользователя по email и паролю
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   signInRequest body requests.SignInRequest true "Данные для входа"
// @Success 200 {object} response.SignInResponse "Tokens"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации или неверные учетные данные"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/sign-in [post]
func (a *AuthController) SignIn(ctx echo.Context) error {
	a.log.Infof("(AuthController.SignIn)")
	var req requests.SignInRequest
	if err := a.decodeRequest(ctx, &req); err != nil {
		a.log.Debugf("Failed to validate request SignIn: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	dbUser, err := a.userRepository.GetByEmail(context.Background(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Internal server error: %v", err)})
	}

	hashedPassword, err := values.NewPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Internal server error by check password: %v", err)})
	}

	if err := ComparePasswords(*hashedPassword, req.Password); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("%v", err)})
	}

	deviceUuid, _ := uuid.NewV7()
	deviceId := deviceUuid.String()
	accessToken, refreshToken, err := a.jwtManager.GenerateTokens(context.Background(), dbUser.Id, deviceId, req.Email, dbUser.Role)
	if err != nil {
		a.log.Debugf("(GenerateTokens) Failed to generate tokens: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Internal server error: %v", err)})
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

// SignOut godoc
// @Summary Выход пользователя
// @Description Завершение сессии пользователя с аннулированием токенов
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   signOutRequest body requests.SignOutRequest true "Данные для выхода"
// @Success 200 {object} response.SignOutResponse "Пустой ответ при успешном завершении"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации или неверные данные"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/sign-out [post]
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
