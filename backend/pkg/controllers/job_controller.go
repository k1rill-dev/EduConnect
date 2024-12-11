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
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type JobController struct {
	log            logger.Logger
	cfg            *config.Config
	validate       *validator.Validate
	jwtManager     jwt.JwtManager
	jobRepository  repository.JobRepository
	userRepository repository.UserRepository
}

func NewJobController(log logger.Logger, cfg *config.Config, jobRepository repository.JobRepository,
	validator *validator.Validate, jwtManager jwt.JwtManager, userRepository repository.UserRepository) *JobController {
	return &JobController{log: log, cfg: cfg, jobRepository: jobRepository, validate: validator, jwtManager: jwtManager,
		userRepository: userRepository}
}

// CreateJob godoc
// @Summary Создать вакансию
// @Description Создает новую вакансию
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   job body requests.CreateJobRequest true "Данные для создания вакансии"
// @Success 201 {object} response.SuccessResponse "Успешно создано"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/jobs [post]
func (c *JobController) CreateJob(ctx echo.Context) error {
	c.log.Infof("(JobController.CreateJob)")
	var request requests.CreateJobRequest
	if err := ctx.Bind(&request); err != nil {
		c.log.Debugf("Failed to bind request: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request payload"})
	}
	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)

	accountId := accountClaims["sub"].(string)
	user, err := c.userRepository.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Error: fmt.Sprintf("Unauthorized: %v", err)})
	}
	if err := c.validate.Struct(request); err != nil {
		c.log.Debugf("Validation error: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: fmt.Sprintf("Validation error: %v", err)})
	}
	id, err := uuid.NewV7()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
	}
	job := model.NewJob(id.String(), user.Id, request.Title, request.Description, request.Location)
	if err := c.jobRepository.Create(context.Background(), job); err != nil {
		c.log.Error("Failed to create request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to create request"})
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse{Message: "Job created successfully"})
}

// UpdateJob godoc
// @Summary Обновить вакансию
// @Description Обновляет данные существующей вакансии
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   jobId path string true "ID вакансии"
// @Param   updateJob body requests.UpdateJobRequest true "Данные для обновления вакансии"
// @Success 200 {object} response.SuccessResponse "Успешно обновлено"
// @Failure 400 {object} response.ErrorResponse "Ошибка валидации"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/jobs/{jobId} [put]
func (c *JobController) UpdateJob(ctx echo.Context) error {
	c.log.Infof("(JobController.UpdateJob)")
	jobId := ctx.Param("jobId")
	var req requests.UpdateJobRequest
	if err := ctx.Bind(&req); err != nil {
		c.log.Debugf("Failed to bind update data: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request payload"})
	}

	accountClaims := (ctx.Get("claims")).(jwt5.MapClaims)

	_ = accountClaims["sub"].(string)
	updatedData := repository.UpdateJob{
		Id:          jobId,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
	}
	if err := c.jobRepository.Update(context.Background(), &updatedData); err != nil {
		c.log.Error("Failed to update job: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to update job"})
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse{Message: "Job updated successfully"})
}

// GetJobById godoc
// @Summary Получить вакансию по ID
// @Description Возвращает вакансию по ее ID
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   jobId path string true "ID вакансии"
// @Success 200 {object} model.Job "Информация о вакансии"
// @Failure 404 {object} response.ErrorResponse "Вакансия не найдена"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/jobs/{jobId} [get]
func (c *JobController) GetJobById(ctx echo.Context) error {
	c.log.Infof("(JobController.GetJobById)")
	jobId := ctx.Param("jobId")
	job, err := c.jobRepository.GetById(context.Background(), jobId)
	if err != nil {
		c.log.Error("Failed to get job: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to retrieve job"})
	}
	if job == nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{Error: "Job not found"})
	}

	return ctx.JSON(http.StatusOK, job)
}

// SearchJobs godoc
// @Summary Поиск вакансий
// @Description Поиск вакансий по названию
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   title query string true "Название вакансии"
// @Param   page query int false "Номер страницы"
// @Param   limit query int false "Количество записей на странице"
// @Success 200 {object} []model.Job "Список вакансий"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/jobs/search [get]
func (c *JobController) SearchJobs(ctx echo.Context) error {
	c.log.Infof("(JobController.SearchJobs)")
	title := ctx.QueryParam("title")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	jobs, err := c.jobRepository.Search(context.Background(), title, page, limit)
	if err != nil {
		c.log.Error("Failed to search jobs: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to search jobs"})
	}

	return ctx.JSON(http.StatusOK, jobs)
}

// GetJobsByFilters godoc
// @Summary Фильтр вакансий
// @Description Возвращает список вакансий, соответствующих заданным фильтрам
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   filters body repository.JobFilters true "Фильтры вакансий"
// @Param   page query int false "Номер страницы"
// @Param   limit query int false "Количество записей на странице"
// @Success 200 {object} []model.Job "Список вакансий"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/jobs/filter [post]
func (c *JobController) GetJobsByFilters(ctx echo.Context) error {
	c.log.Infof("(JobController.GetJobsByFilters)")
	var filters repository.JobFilters
	if err := ctx.Bind(&filters); err != nil {
		c.log.Debugf("Failed to bind filters: %v", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request payload"})
	}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	jobs, err := c.jobRepository.GetByFilters(context.Background(), &filters, page, limit)
	if err != nil {
		c.log.Error("Failed to get jobs by filters: %v", err)
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to get jobs by filters"})
	}

	return ctx.JSON(http.StatusOK, jobs)
}
