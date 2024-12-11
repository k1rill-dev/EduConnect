package repo

import (
	"EduConnect/internal/model"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioRepositoryMongo struct {
	log logger.Logger
	cfg *config.Config
	db  *mongo.Collection
}

func NewPortfolioRepositoryMongo(log logger.Logger, cfg *config.Config, db *mongo.Client) *PortfolioRepositoryMongo {
	collection := db.Database(cfg.Mongo.Db).Collection("portfolio")
	return &PortfolioRepositoryMongo{log: log, cfg: cfg, db: collection}
}
func (r *PortfolioRepositoryMongo) Create(ctx context.Context, portfolio *model.Portfolio) error {
	_, err := r.db.InsertOne(ctx, portfolio)
	return err
}

func (r *PortfolioRepositoryMongo) AddItems(ctx context.Context, studentId string, items []model.PortfolioItems) error {
	_, err := r.db.UpdateOne(
		ctx,
		bson.M{"student_id": studentId},
		bson.M{"$push": bson.M{"items": bson.M{"$each": items}}},
	)
	return err
}

func (r *PortfolioRepositoryMongo) GetByStudentId(ctx context.Context, studentId string) (*model.Portfolio, error) {
	var portfolio model.Portfolio
	err := r.db.FindOne(ctx, bson.M{"student_id": studentId}).Decode(&portfolio)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}
	return &portfolio, nil
}
