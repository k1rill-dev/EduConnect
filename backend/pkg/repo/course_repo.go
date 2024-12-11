package repo

import (
	"EduConnect/internal/model"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
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

func (c *CourseMongoRepo) List(ctx context.Context) ([]*model.Course, error) {
	var courses []*model.Course
	cursor, err := c.getCourseCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get courses")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var course model.Course
		if err := cursor.Decode(&course); err != nil {
			return nil, fmt.Errorf("failed to get course from cursor")
		}
		courses = append(courses, &course)
	}

	if err := cursor.Err(); err != nil {
		c.log.Debugf("Cursor error: %v", err)
		return nil, fmt.Errorf("cursor error")
	}
	return courses, nil
}

func (c *CourseMongoRepo) Create(ctx context.Context, course *model.Course) error {
	_, err := c.getCourseCollection().InsertOne(ctx, course, &options.InsertOneOptions{})
	if err != nil && !strings.Contains(err.Error(), "no documents") {
		c.log.Debugf("(CourseMOongoRepo) error: %v", err)
		return err
	}
	return nil
}

func (c *CourseMongoRepo) GetById(ctx context.Context, courseId string) (*model.Course, error) {
	var course model.Course
	err := c.getCourseCollection().FindOne(ctx, bson.M{"_id": courseId}).Decode(&course)
	if err != nil {
		c.log.Debugf("(CourseMongoRepo) error: %v", err)
		return nil, err
	}
	return &course, nil
}

func (c *CourseMongoRepo) SubmitAssignment(ctx context.Context, submission *model.Submission) error {
	return nil
}

func (c *CourseMongoRepo) getCourseCollection() *mongo.Collection {
	return c.mongoClient.Database(c.cfg.Mongo.Db).Collection(c.cfg.MongoCollections.Courses)
}
