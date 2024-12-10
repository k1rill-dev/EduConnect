package repository

import (
	"EduConnect/internal/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, userId string) (*model.User, error)
}
