package repository

import (
	"EduConnect/internal/model"
	"EduConnect/internal/values"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, userId string) (*model.User, error)
	GetByEmail(ctx context.Context, email *values.Email) (*model.User, error)
}
