package repository

import (
	"EduConnect/internal/model"
	"context"
)

type PortfolioRepository interface {
	Create(ctx context.Context, portfolio *model.Portfolio) error
	AddItems(ctx context.Context, studentId string, items []model.PortfolioItems) error
	GetByStudentId(ctx context.Context, studentId string) (*model.Portfolio, error)
}
