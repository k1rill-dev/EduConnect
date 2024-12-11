package repository

import (
	"EduConnect/internal/model"
	"context"
)

type Pagination struct {
	Limit  int `json:"limit"`  // Количество записей на странице
	Offset int `json:"offset"` // Смещение
}

type JobApplicationRepository interface {
	Create(ctx context.Context, application *model.JobApplication) error
	Delete(ctx context.Context, applicationId string) error
	UpdateStatus(ctx context.Context, applicationId string, status string) error

	GetByCompany(ctx context.Context, companyId string) ([]*model.JobApplication, error)
	GetByStudent(ctx context.Context, studentId string) ([]*model.JobApplication, error)
}
