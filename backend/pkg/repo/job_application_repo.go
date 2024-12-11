package repo

import (
	"EduConnect/internal/model"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobApplicationRepo struct {
	log logger.Logger
	cfg *config.Config
	db  *mongo.Collection
}

func NewJobApplicationRepo(log logger.Logger, cfg *config.Config, db *mongo.Client) *JobApplicationRepo {
	collection := db.Database(cfg.Mongo.Db).Collection("jobs_application")
	return &JobApplicationRepo{log: log, cfg: cfg, db: collection}
}
func (r *JobApplicationRepo) Create(ctx context.Context, application *model.JobApplication) error {
	_, err := r.db.InsertOne(ctx, application)
	return err
}

func (r *JobApplicationRepo) Delete(ctx context.Context, applicationId string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": applicationId})
	return err
}

func (r *JobApplicationRepo) UpdateStatus(ctx context.Context, applicationId string, status string) error {
	_, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": applicationId},
		bson.M{"$set": bson.M{"status": status}},
	)
	return err
}

func (r *JobApplicationRepo) GetByCompany(ctx context.Context, companyId string) ([]*model.JobApplication, error) {
	cursor, err := r.db.Find(ctx, bson.M{"company_id": companyId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var applications []*model.JobApplication
	for cursor.Next(ctx) {
		var application model.JobApplication
		if err := cursor.Decode(&application); err != nil {
			return nil, err
		}
		applications = append(applications, &application)
	}
	return applications, nil
}

func (r *JobApplicationRepo) GetByStudent(ctx context.Context, studentId string) ([]*model.JobApplication, error) {
	cursor, err := r.db.Find(ctx, bson.M{"student_id": studentId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var applications []*model.JobApplication
	for cursor.Next(ctx) {
		var application model.JobApplication
		if err := cursor.Decode(&application); err != nil {
			return nil, err
		}
		applications = append(applications, &application)
	}
	return applications, nil
}
