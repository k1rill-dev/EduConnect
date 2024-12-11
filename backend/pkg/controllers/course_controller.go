package controllers

import (
	"EduConnect/internal/model"
	"EduConnect/internal/repository"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"EduConnect/pkg/requests"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

const pathToPdf = "storage/assignments/"

type CourseController struct {
	log        logger.Logger
	cfg        *config.Config
	validate   *validator.Validate
	courseRepo repository.CourseRepository
}

func NewCourseController(log logger.Logger, cfg *config.Config, validate *validator.Validate, courseRepo repository.CourseRepository) *CourseController {
	return &CourseController{log: log, cfg: cfg, validate: validate, courseRepo: courseRepo}
}

func (c *CourseController) generateAssignmentFilename() string {
	uuid, _ := uuid.NewV7()
	return uuid.String()
}

func (c *CourseController) CreateCourse(ctx echo.Context) error {
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	teacherID := ctx.FormValue("teacher_id")
	layout := time.RFC3339
	startDate, _ := time.Parse(layout, ctx.FormValue("start_date"))
	endDate, _ := time.Parse(layout, ctx.FormValue("end_date"))

	createCourseReq := &requests.CreateCourseRequest{
		Title:       title,
		Description: description,
		TeacherId:   teacherID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := c.validate.Struct(createCourseReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Ошибка валидации: %v", err),
		})
	}

	topicsFile, err := ctx.FormFile("topics")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Не удалось получить файл с темами",
		})
	}

	topicsData, err := topicsFile.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Ошибка открытия файла с темами",
		})
	}
	defer topicsData.Close()

	var topicsReq []requests.TopicRequest
	if err := json.NewDecoder(topicsData).Decode(&topicsReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Ошибка парсинга JSON с темами",
		})
	}

	topics := []*model.Topic{}
	for _, topic := range topicsReq {
		topicModel := &model.Topic{}
		assignments := []*model.Assignment{}

		for _, assignment := range topic.Assignments {
			filePath, err := c.savePdfFile(ctx, assignment.TheoryFile)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}

			assignmentModel := &model.Assignment{
				Title:          assignment.Title,
				TheoryFile:     filePath,
				AdditionalInfo: assignment.AdditionalInfo,
			}

			assignments = append(assignments, assignmentModel)
		}
		topicModel.Title = topic.Title
		topicModel.Assignments = assignments
		topics = append(topics, topicModel)
	}

	courseUuid, _ := uuid.NewV7()
	courseId := courseUuid.String()
	course := &model.Course{
		Id:          courseId,
		Title:       createCourseReq.Title,
		Description: createCourseReq.Description,
		TeacherId:   createCourseReq.TeacherId,
		StartDate:   createCourseReq.StartDate,
		EndDate:     createCourseReq.EndDate,
		Topics:      topics,
		CreatedAt:   time.Now(),
	}

	if err := c.courseRepo.Create(context.Background(), course); err != nil {
		c.log.Error("Error by save course: ", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error by save course into bd",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Курс успешно создан",
		"course":  course,
	})
}

func (c *CourseController) savePdfFile(ctx echo.Context, theoryFile string) (string, error) {
	if _, err := os.Stat(pathToPdf); os.IsNotExist(err) {
		if err := os.MkdirAll(pathToPdf, os.ModePerm); err != nil {
			c.log.Error("Ошибка при создании папки для файлов: ", err)
			return "", fmt.Errorf("ошибка при создании папки для файлов: %v", err)
		}
	}
	pdfFile, err := ctx.FormFile(theoryFile)
	if err != nil {
		return "", ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Файл для задачи %s не найден", theoryFile),
		})
	}
	filename := c.generateAssignmentFilename()
	filePath := fmt.Sprintf("%s%s.pdf", pathToPdf, filename)
	fileUrl := fmt.Sprintf("http://localhost%s/api/%s", c.cfg.Http.Port, filePath)
	src, err := pdfFile.Open()
	if err != nil {
		c.log.Error("Error by open pdf: ", err)
		return "", fmt.Errorf("ошибка открытия файла PDF")
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		c.log.Error("Error by save pdf: ", err)
		return "", fmt.Errorf("ошибка сохранения файла PDF")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		c.log.Error("Error by copy pdf: ", err)
		return "", fmt.Errorf("ошибка копирования файла PDF")
	}

	return fileUrl, nil
}
