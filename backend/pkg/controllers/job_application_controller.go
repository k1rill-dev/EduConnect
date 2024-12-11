package controllers

import (
	"EduConnect/internal/model"
	"EduConnect/internal/repository"
	"EduConnect/pkg/config"
	"EduConnect/pkg/jwt"
	"EduConnect/pkg/logger"
	"EduConnect/pkg/requests"
	"EduConnect/pkg/response"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ApplicationController struct {
	log                   logger.Logger
	cfg                   *config.Config
	validate              *validator.Validate
	jwtManager            jwt.JwtManager
	applicationRepository repository.JobApplicationRepository
	userRepository        repository.UserRepository
}

func NewApplicationController(log logger.Logger, cfg *config.Config, applicationRepository repository.JobApplicationRepository,
	validator *validator.Validate, jwtManager jwt.JwtManager, userRepository repository.UserRepository) *ApplicationController {
	return &ApplicationController{
		log: log, cfg: cfg,
		applicationRepository: applicationRepository,
		validate:              validator,
		jwtManager:            jwtManager,
		userRepository:        userRepository,
	}
}

// CreateApplication godoc
// @Summary Создать отклик
// @Description Создает новый отклик на вакансию
// @Tags applications
// @Accept  json
// @Produce  json
// @Param   application body requests.CreateJobApplication true "Данные отклика"
// @Success 201 {object} response.SuccessResponse "Успешно создано"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/applications [post]
func (c *ApplicationController) CreateApplication(ctx echo.Context) error {
	c.log.Infof("(ApplicationController.CreateApplication)")
	var req requests.CreateJobApplication
	if err := ctx.Bind(&req); err != nil {
		c.log.Debugf("Failed to bind request: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request payload"})
	}

	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)
	accountId := accountClaims["sub"].(string)

	// Получаем данные пользователя для проверки роли
	user, err := c.userRepository.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: "Unauthorized"})
	}

	if user.Role != "student" {
		return ctx.JSON(http.StatusForbidden, response.ErrorResponse{Error: "Only students can create applications"})
	}
	id, err := uuid.NewV7()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
	}
	data := model.NewJobApplication(id.String(), req.CompanyId, req.StudentId, "unread")
	if err := c.applicationRepository.Create(context.Background(), data); err != nil {
		c.log.Error("Failed to create req: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to create req"})
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse{Message: "Application created successfully"})
}

// UpdateApplicationStatus godoc
// @Summary Обновить статус отклика
// @Description Обновляет статус существующего отклика
// @Tags applications
// @Accept  json
// @Produce  json
// @Param   applicationId path string true "ID отклика"
// @Param   status query string true "Новый статус"
// @Success 200 {object} response.SuccessResponse "Успешно обновлено"
// @Failure 400 {object} response.ErrorResponse "Ошибка запроса"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/applications/{applicationId}/status [put]
func (c *ApplicationController) UpdateApplicationStatus(ctx echo.Context) error {
	c.log.Infof("(ApplicationController.UpdateApplicationStatus)")
	applicationId := ctx.Param("applicationId")
	newStatus := ctx.QueryParam("status")

	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)
	accountId := accountClaims["sub"].(string)

	_, err := c.userRepository.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: "Unauthorized"})
	}

	if err := c.applicationRepository.UpdateStatus(context.Background(), applicationId, newStatus); err != nil {
		c.log.Error("Failed to update application status: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to update application status"})
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse{Message: "Application status updated successfully"})
}

// DeleteApplication godoc
// @Summary Удалить отклик
// @Description Удаляет отклик
// @Tags applications
// @Accept  json
// @Produce  json
// @Param   applicationId path string true "ID отклика"
// @Success 200 {object} response.SuccessResponse "Успешно удалено"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/applications/{applicationId} [delete]
func (c *ApplicationController) DeleteApplication(ctx echo.Context) error {
	c.log.Infof("(ApplicationController.DeleteApplication)")
	applicationId := ctx.Param("applicationId")

	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)
	accountId := accountClaims["sub"].(string)

	_, err := c.userRepository.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: "Unauthorized"})
	}

	if err := c.applicationRepository.Delete(context.Background(), applicationId); err != nil {
		c.log.Error("Failed to delete application: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse{Message: "Application deleted successfully"})
}
