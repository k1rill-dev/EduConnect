package repo

import (
	"EduConnect/internal/model"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseMongoRepo struct {
	log         logger.Logger
	cfg         *config.Config
	mongoClient *mongo.Client
}

func NewCourseMongoRepo(log logger.Logger, cfg *config.Config, mongoClient *mongo.Client) *CourseMongoRepo {
	return &CourseMongoRepo{log: log, cfg: cfg, mongoClient: mongoClient}
}

func (c *CourseMongoRepo) Create(ctx context.Context, course *model.Course) error {
	_, err := c.getCourseCollection().InsertOne(ctx, course, &options.InsertOneOptions{})
	if err != nil && !strings.Contains(err.Error(), "no documents") {
		c.log.Debugf("(CourseMOongoRepo) error: %v", err)
		return err
	}
	return nil
}

func (c *CourseMongoRepo) getCourseCollection() *mongo.Collection {
	return c.mongoClient.Database(c.cfg.Mongo.Db).Collection(c.cfg.MongoCollections.Courses)
}
