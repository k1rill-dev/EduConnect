package repository

import (
	"EduConnect/internal/model"
	"EduConnect/internal/values"
	"context"
)

type UpdateUserRequest struct {
	Id        string        `bson:"_id"`
	Email     *values.Email `json:"email" validate:"required,email"`
	FirstName string        `bson:"first_name" validate:"required"`
	Surname   string        `bson:"surname" validate:"required"`
	Picture   string        `json:"picture" validate:"required"`
	Bio       string        `json:"bio" validate:"required"`
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *UpdateUserRequest) error
	GetById(ctx context.Context, userId string) (*model.User, error)
	GetByEmail(ctx context.Context, email *values.Email) (*model.User, error)
}
