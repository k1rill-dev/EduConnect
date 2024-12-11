package repository

import (
	"EduConnect/internal/model"
	"context"
)

type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
	GetById(ctx context.Context, courseId string) (*model.Course, error)
	SubmitAssignment(ctx context.Context, submission *model.Submission) error
	List(ctx context.Context) ([]*model.Course, error)
}
