package repo

import (
	"EduConnect/internal/model"
	"EduConnect/internal/repository"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type JobRepo struct {
	log logger.Logger
	cfg *config.Config
	db  *mongo.Collection
}

func NewJobRepo(log logger.Logger, cfg *config.Config, db *mongo.Client) *JobRepo {
	collection := db.Database(cfg.Mongo.Db).Collection("jobs")
	return &JobRepo{log: log, cfg: cfg, db: collection}
}

func (repo *JobRepo) Create(ctx context.Context, job *model.Job) error {
	job.Id = primitive.NewObjectID().Hex()
	job.CreatedAt = time.Now()
	_, err := repo.db.InsertOne(ctx, job)
	return err
}

func (repo *JobRepo) Update(ctx context.Context, job *repository.UpdateJob) error {
	filter := bson.M{"_id": job.Id}
	update := bson.M{
		"$set": bson.M{
			"title":       job.Title,
			"description": job.Description,
			"location":    job.Location,
		},
	}
	_, err := repo.db.UpdateOne(ctx, filter, update)
	return err
}

func (repo *JobRepo) GetById(ctx context.Context, jobId string) (*model.Job, error) {
	var job model.Job
	err := repo.db.FindOne(ctx, bson.M{"_id": jobId}).Decode(&job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (repo *JobRepo) Search(ctx context.Context, title string, page, limit int) (*[]model.Job, error) {
	filter := bson.M{"title": bson.M{"$regex": title, "$options": "i"}}
	return repo.paginatedFind(ctx, filter, page, limit)
}

func (repo *JobRepo) GetByFilters(ctx context.Context, filters *repository.JobFilters, page, limit int) (*[]model.Job, error) {
	filter := bson.M{}
	if filters.Location != nil {
		filter["location"] = *filters.Location
	}
	if filters.EmployerId != nil {
		filter["employer_id"] = *filters.EmployerId
	}
	return repo.paginatedFind(ctx, filter, page, limit)
}

func (repo *JobRepo) paginatedFind(ctx context.Context, filter bson.M, page, limit int) (*[]model.Job, error) {
	skip := (page - 1) * limit
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	cursor, err := repo.db.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var jobs []model.Job
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, err
	}
	return &jobs, nil
}
