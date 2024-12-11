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
	_, err := c.getSubmissionCollection().InsertOne(ctx, submission, &options.InsertOneOptions{})
	if err != nil && !strings.Contains(err.Error(), "no documents") {
		c.log.Debugf("(CourseMOongoRepo) error: %v", err)
		return err
	}
	return nil
}

func (c *CourseMongoRepo) GetSubmissionsByStudentId(ctx context.Context, studentId string) ([]*model.Submission, error) {
	submissions := []*model.Submission{}
	cursor, err := c.getSubmissionCollection().Find(ctx, bson.M{"student_id": studentId})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var submission model.Submission
		if err := cursor.Decode(&submission); err != nil {
			return nil, err
		}
		submissions = append(submissions, &submission)
	}

	return submissions, nil
}

func (c *CourseMongoRepo) GetSubmissionsByTeacherId(ctx context.Context, teacherId string) ([]*model.Submission, error) {
	submissions := []*model.Submission{}
	cursor, err := c.getSubmissionCollection().Find(ctx, bson.M{"teacher_id": teacherId})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var submission model.Submission
		if err := cursor.Decode(&submission); err != nil {
			return nil, err
		}
		submissions = append(submissions, &submission)
	}

	return submissions, nil
}

func (c *CourseMongoRepo) EnrollCourse(ctx context.Context, enrollment *model.CourseEnrollment) error {
	_, err := c.getEnrollmentCollection().InsertOne(ctx, enrollment, &options.InsertOneOptions{})
	if err != nil && !strings.Contains(err.Error(), "no documents") {
		c.log.Debugf("(CourseMongoRepo) error: %v", err)
		return err
	}
	return nil
}

func (c *CourseMongoRepo) UpdateSubmission(ctx context.Context, submission *model.Submission) error {
	_, err := c.getSubmissionCollection().UpdateOne(ctx, bson.M{"_id": submission.Id}, submission)
	if err != nil && !strings.Contains(err.Error(), "no documents") {
		c.log.Debugf("(CourseMongoRepo) error: %v", err)
		return err
	}
	return nil
}

func (c *CourseMongoRepo) GetSubmissionById(ctx context.Context, submissionId string) (*model.Submission, error) {
	var course model.Submission
	err := c.getCourseCollection().FindOne(ctx, bson.M{"_id": submissionId}).Decode(&course)
	if err != nil {
		c.log.Debugf("(CourseMongoRepo) error: %v", err)
		return nil, err
	}
	return &course, nil
}

func (c *CourseMongoRepo) getEnrollmentCollection() *mongo.Collection {
	return c.mongoClient.Database(c.cfg.Mongo.Db).Collection(c.cfg.MongoCollections.Enrollments)
}

func (c *CourseMongoRepo) getCourseCollection() *mongo.Collection {
	return c.mongoClient.Database(c.cfg.Mongo.Db).Collection(c.cfg.MongoCollections.Courses)
}

func (c *CourseMongoRepo) getSubmissionCollection() *mongo.Collection {
	return c.mongoClient.Database(c.cfg.Mongo.Db).Collection(c.cfg.MongoCollections.Submissions)
}
