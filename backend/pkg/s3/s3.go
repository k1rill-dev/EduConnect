package s3

import (
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const storagePath = "./storage"

type S3Storage struct {
	cfg         *config.Config
	log         logger.Logger
	mongoClient *mongo.Client
}

type FileInfo struct {
	ID        string `bson:"_id,omitempty"`
	FileName  string `bson:"filename"`
	UploadURL string `bson:"upload_url"`
}

func NewS3Storage(log logger.Logger, cfg *config.Config, mongoClient *mongo.Client) (*S3Storage, error) {
	err := os.MkdirAll(storagePath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &S3Storage{
		log:         log,
		cfg:         cfg,
		mongoClient: mongoClient,
	}, nil
}

func (s *S3Storage) generateFileName() string {
	uuid, _ := uuid.NewV7()
	return uuid.String()
}

func (s *S3Storage) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Не удалось получить файл")
	}
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось открыть файл")
	}
	defer src.Close()

	fileID := fmt.Sprintf("%s.jpg", s.generateFileName())
	filePath := filepath.Join(storagePath, fileID)

	dst, err := os.Create(filePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось создать файл на сервере")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при копировании файла")
	}

	downloadURL := fmt.Sprintf("http://localhost:8083/files/%s", fileID)

	defer s.mongoClient.Disconnect(context.TODO())

	fileInfo := FileInfo{
		ID:        fileID,
		FileName:  file.Filename,
		UploadURL: downloadURL,
	}

	_, err = s.getS3Collection().InsertOne(context.TODO(), fileInfo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось сохранить информацию о файле в базу данных")
	}

	return c.String(http.StatusOK, downloadURL)
}

func (s *S3Storage) GetFileLink(c echo.Context) error {
	fileID := c.Param("id")

	defer s.mongoClient.Disconnect(context.TODO())

	var fileInfo FileInfo
	err := s.getS3Collection().FindOne(context.TODO(), bson.M{"_id": fileID}).Decode(&fileInfo)
	if err != nil {
		return c.String(http.StatusNotFound, "Файл не найден")
	}

	return c.String(http.StatusOK, fileInfo.UploadURL)
}

func (s *S3Storage) DownloadFile(c echo.Context) error {
	fileID := c.Param("id")

	filePath := filepath.Join(storagePath, fileID)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.String(http.StatusNotFound, "Файл не найден")
	}

	return c.File(filePath)
}

func (s *S3Storage) DeleteFile(c echo.Context) error {
	fileID := c.Param("id")

	defer s.mongoClient.Disconnect(context.TODO())

	_, err := s.getS3Collection().DeleteOne(context.TODO(), bson.M{"_id": fileID})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при удалении информации о файле из базы данных")
	}

	filePath := filepath.Join(storagePath, fileID)

	err = os.Remove(filePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при удалении файла")
	}

	return c.String(http.StatusOK, fmt.Sprintf("Файл удален: %s", fileID))
}

func (s *S3Storage) getS3Collection() *mongo.Collection {
	return s.mongoClient.Database(s.cfg.Mongo.Db).Collection(s.cfg.MongoCollections.S3)
}
