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
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const pathToPdf = "storage/assignments/"

type CourseController struct {
	log        logger.Logger
	cfg        *config.Config
	validate   *validator.Validate
	courseRepo repository.CourseRepository
	userRepo   repository.UserRepository
}

func NewCourseController(log logger.Logger, cfg *config.Config, validate *validator.Validate, courseRepo repository.CourseRepository, userRepo repository.UserRepository) *CourseController {
	return &CourseController{log: log, cfg: cfg, validate: validate, courseRepo: courseRepo, userRepo: userRepo}
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

	topicsStr := ctx.FormValue("topics")
	if topicsStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Поле topics не передано",
		})
	}

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

	// topicsFile, err := ctx.FormFile("topics")
	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, map[string]string{
	// 		"error": "Не удалось получить файл с темами",
	// 	})
	// }

	// topicsData, err := topicsFile.Open()
	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, map[string]string{
	// 		"error": "Ошибка открытия файла с темами",
	// 	})
	// }
	// defer topicsData.Close()

	var topicsReq []requests.TopicRequest
	if err := json.Unmarshal([]byte(topicsStr), &topicsReq); err != nil {
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

func (c *CourseController) GetCourses(ctx echo.Context) error {
	courses, err := c.courseRepo.List(context.Background())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "course",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"courses": courses,
	})
}

func (c *CourseController) GetCourseById(ctx echo.Context) error {
	var req requests.GetCourseByIdRequest
	if err := c.decodeRequest(ctx, &req); err != nil {
		c.log.Debugf("Failed to decode request GetCourseById: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}
	course, err := c.courseRepo.GetById(context.Background(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Error by get course by id: %v", err)})
	}

	return ctx.JSON(http.StatusOK, course)
}

func (c *CourseController) SubmitAssignment(ctx echo.Context) error {
	// var req requests.SubmitAssignmentRequest
	// if err := c.decodeRequest(ctx, &req); err != nil {
	// 	c.log.Debugf("Failed to decode request GetCourseById: %v", err)
	// 	return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	// }
	accountClaims := (ctx.Get("claims")).(jwt.MapClaims)
	accountId := accountClaims["sub"].(string)

	account, err := c.userRepo.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	if account.Role != "student" {
		return ctx.JSON(http.StatusNetworkAuthenticationRequired, "Role is not student")
	}

	topic := ctx.FormValue("topic")
	assignment := ctx.FormValue("assignment")
	courseId := ctx.FormValue("course_id")
	teacherId := ctx.FormValue("teacher_id")

	filePath, err := c.saveFile(ctx, "submission")
	if err != nil {
		c.log.Debugf("Error: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error by save pdf",
		})
	}

	req := requests.SubmitAssignmentRequest{
		Topic:      topic,
		Assignment: assignment,
		Submission: filePath,
		CourseId:   courseId,
		TeacherId:  teacherId,
	}

	submissionUuid, _ := uuid.NewV7()
	submissionId := submissionUuid.String()
	submission := &model.Submission{
		Id:          submissionId,
		Topic:       req.Topic,
		Assignment:  req.Assignment,
		Submission:  req.Submission,
		SubmittedAt: time.Now(),
		CourseId:    req.CourseId,
		StudentId:   accountId,
		TeacherId:   req.TeacherId,
	}

	if err := c.courseRepo.SubmitAssignment(context.Background(), submission); err != nil {
		c.log.Error("Error by save submit assignment: ", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error by save submit assignment",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{})
}

func (c *CourseController) GetSubmissionsByTeacherId(ctx echo.Context) error {
	var req requests.GetSubmissionsByTeacherIdRequest
	if err := c.decodeRequest(ctx, &req); err != nil {
		c.log.Debugf("Failed to decode request GetCourseById: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	submissions, err := c.courseRepo.GetSubmissionsByTeacherId(context.Background(), req.TeacherId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error by get submissions",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"submissions": submissions,
	})
}

func (c *CourseController) GetSubmissionById(ctx echo.Context) error {
	var req requests.GetSubmissionById
	if err := c.decodeRequest(ctx, &req); err != nil {
		c.log.Debugf("Failed to decode request GetCourseById: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	submission, err := c.courseRepo.GetById(context.Background(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"submission": submission,
	})
}

func (c *CourseController) GetByStudentId(ctx echo.Context) error {
	accountClaims := (ctx.Get("claims")).(jwt.MapClaims)
	accountId := accountClaims["sub"].(string)

	// account, err := c.userRepo.GetById(context.Background(), accountId)
	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	// }

	if user, _ := c.userRepo.GetById(context.Background(), accountId); user == nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "User doesnt exist"})
	}

	submissions, err := c.courseRepo.GetSubmissionsByStudentId(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"submissions": submissions,
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
	fileUrl := fmt.Sprintf("http://localhost%s/api/file/%s.pdf", c.cfg.Http.Port, filename)
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

func (a *CourseController) decodeRequest(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return fmt.Errorf("invalid request")
	}

	if err := a.validate.Struct(i); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}

func (c *CourseController) EnrollCourse(ctx echo.Context) error {
	accountClaims := (ctx.Get("claims")).(jwt.MapClaims)
	accountId := accountClaims["sub"].(string)

	account, err := c.userRepo.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	if account.Role != "student" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "role is not student"})
	}

	var req requests.EnrollStudentRequest
	if err := c.decodeRequest(ctx, &req); err != nil {
		c.log.Debugf("Failed to decode request GetCourseById: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	enrollmentCourseUuid, _ := uuid.NewV7()
	enrollmentId := enrollmentCourseUuid.String()
	enrollment := &model.CourseEnrollment{
		Id:         enrollmentId,
		StudentId:  accountId,
		CourseId:   req.CourseId,
		EnrolledAt: time.Now(),
	}

	if err := c.courseRepo.EnrollCourse(context.Background(), enrollment); err != nil {
		c.log.Error("Failed to save enrollment: ", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to save enrollent",
		})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c *CourseController) EvaluateStudent(ctx echo.Context) error {
	accountClaims := (ctx.Get("claims")).(jwt.MapClaims)
	accountId := accountClaims["sub"].(string)

	account, err := c.userRepo.GetById(context.Background(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	if account.Role != "teacher" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "role is not teacher"})
	}

	var req requests.EvaluateStudentRequest
	if err := c.decodeRequest(ctx, &req); err != nil {
		c.log.Debugf("Failed to decode request GetCourseById: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Validation error: %v", err)})
	}

	submission, err := c.courseRepo.GetSubmissionById(context.Background(), req.SubmissionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	submission.Grade = req.Grade

	if err := c.courseRepo.UpdateSubmission(context.Background(), submission); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update submission",
		})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c *CourseController) saveFile(ctx echo.Context, formKey string) (string, error) {
	if _, err := os.Stat(pathToPdf); os.IsNotExist(err) {
		if err := os.MkdirAll(pathToPdf, os.ModePerm); err != nil {
			c.log.Error("Error creating directory: ", err)
			return "", fmt.Errorf("error creating directory: %v", err)
		}
	}

	fileHeader, err := ctx.FormFile(formKey)
	if err != nil {
		return "", fmt.Errorf("file not found: %s", formKey)
	}

	src, err := fileHeader.Open()
	if err != nil {
		c.log.Error("Error opening file: ", err)
		return "", fmt.Errorf("error opening file")
	}
	defer src.Close()

	allowedExtensions := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
		".txt":  true,
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	ext := filepath.Ext(fileHeader.Filename)
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("file type '%s' is not allowed", ext)
	}

	filename := fmt.Sprintf("%s%s", c.generateAssignmentFilename(), ext)
	filePath := filepath.Join(pathToPdf, filename)
	fileUrl := fmt.Sprintf("http://localhost%s/api/file/%s", c.cfg.Http.Port, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		c.log.Error("Error saving file: ", err)
		return "", fmt.Errorf("error saving file")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		c.log.Error("Error copying file: ", err)
		return "", fmt.Errorf("error copying file")
	}

	return fileUrl, nil
}
