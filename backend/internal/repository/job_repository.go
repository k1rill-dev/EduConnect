package repository

import (
	"EduConnect/internal/model"
	"context"
)

type UpdateJob struct {
	Id          string `bson:"_id,omitempty"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Location    string `bson:"location"`
}

type JobFilters struct {
	Location   *string `bson:"location"`
	EmployerId *string `bson:"employer_id"`
}

type JobRepository interface {
	Create(ctx context.Context, job *model.Job) error
	Update(ctx context.Context, job *UpdateJob) error
	GetById(ctx context.Context, jobId string) (*model.Job, error)
	Search(ctx context.Context, title string, page, limit int) (*[]model.Job, error)
	GetByFilters(ctx context.Context, filters *JobFilters, page, limit int) (*[]model.Job, error)
}
