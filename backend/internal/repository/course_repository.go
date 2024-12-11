package repository

import (
	"EduConnect/internal/model"
	"context"
)

type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
}
