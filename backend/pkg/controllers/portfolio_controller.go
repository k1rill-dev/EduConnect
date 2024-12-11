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

type PortfolioController struct {
	log                 logger.Logger
	cfg                 *config.Config
	validate            *validator.Validate
	jwtManager          jwt.JwtManager
	portfolioRepository repository.PortfolioRepository
	userRepository      repository.UserRepository
}

func NewPortfolioController(log logger.Logger, cfg *config.Config, portfolioRepository repository.PortfolioRepository,
	validator *validator.Validate, jwtManager jwt.JwtManager, userRepository repository.UserRepository) *PortfolioController {
	return &PortfolioController{
		log:                 log,
		cfg:                 cfg,
		portfolioRepository: portfolioRepository,
		validate:            validator,
		jwtManager:          jwtManager,
		userRepository:      userRepository,
	}
}

// CreatePortfolio godoc
// @Summary Создать портфолио
// @Description Создает новое портфолио
// @Tags portfolio
// @Accept  json
// @Produce  json
// @Param   portfolio body requests.CreatePortfolioRequest true "Данные портфолио"
// @Success 201 {object} response.SuccessResponse "Успешно создано"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/portfolios [post]
func (c *PortfolioController) CreatePortfolio(ctx echo.Context) error {
	c.log.Infof("(PortfolioController.CreatePortfolio)")
	var req requests.CreatePortfolioRequest
	if err := ctx.Bind(&req); err != nil {
		c.log.Debugf("Failed to bind request: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request payload"})
	}

	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)
	accountId := accountClaims["sub"].(string)

	// Проверяем роль пользователя
	user, err := c.userRepository.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: "Unauthorized"})
	}

	if user.Role != "student" {
		return ctx.JSON(http.StatusForbidden, response.ErrorResponse{Error: "Only students can create portfolios"})
	}
	id, err := uuid.NewV7()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
	}
	data := model.NewPortfolio(id.String(), user.Id, req.Items)
	if err := c.portfolioRepository.Create(context.Background(), data); err != nil {
		c.log.Error("Failed to create req: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to create req"})
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse{Message: "Portfolio created successfully"})
}

// AddPortfolioItems godoc
// @Summary Добавить элементы в портфолио
// @Description Добавляет элементы в существующее портфолио
// @Tags portfolio
// @Accept  json
// @Produce  json
// @Param   studentId path string true "ID студента"
// @Param   items body []model.PortfolioItems true "Список элементов портфолио"
// @Success 200 {object} response.SuccessResponse "Элементы добавлены"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/portfolios/{portfolioId}/items [post]
func (c *PortfolioController) AddPortfolioItems(ctx echo.Context) error {
	c.log.Infof("(PortfolioController.AddPortfolioItems)")
	_ = ctx.Param("studentId")
	var items []model.PortfolioItems
	if err := ctx.Bind(&items); err != nil {
		c.log.Debugf("Failed to bind request: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request payload"})
	}
	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)
	accountId := accountClaims["sub"].(string)
	user, err := c.userRepository.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: "Unauthorized"})
	}
	if user.Role != "student" {
		return ctx.JSON(http.StatusForbidden, response.ErrorResponse{Error: "Only students can add portfolio items"})
	}
	if err := c.portfolioRepository.AddItems(context.Background(), accountId, items); err != nil {
		c.log.Error("Failed to add items to portfolio: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to add items to portfolio"})
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse{Message: "Items added to portfolio successfully"})
}

// GetPortfolioByStudent godoc
// @Summary Получить портфолио студента
// @Description Возвращает портфолио по ID студента
// @Tags portfolio
// @Accept  json
// @Produce  json
// @Param   studentId path string true "ID студента"
// @Success 200 {object} model.Portfolio "Портфолио"
// @Failure 404 {object} response.ErrorResponse "Портфолио не найдено"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/portfolios/student/{studentId} [get]
func (c *PortfolioController) GetPortfolioByStudent(ctx echo.Context) error {
	c.log.Infof("(PortfolioController.GetPortfolioByStudent)")
	studentId := ctx.Param("studentId")

	portfolio, err := c.portfolioRepository.GetByStudentId(context.Background(), studentId)
	if err != nil {
		c.log.Error("Failed to get portfolio: %v", err)
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{Error: err.Error()})
	}
	if portfolio == nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{Error: "Portfolio not found"})
	}

	return ctx.JSON(http.StatusOK, portfolio)
}
