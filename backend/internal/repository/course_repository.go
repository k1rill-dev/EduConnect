package repository

import (
	"EduConnect/internal/model"
	"context"
)

type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
	GetById(ctx context.Context, courseId string) (*model.Course, error)
	SubmitAssignment(ctx context.Context, submission *model.Submission) error
	GetSubmissionsByTeacherId(ctx context.Context, teacherId string) ([]*model.Submission, error)
	GetSubmissionsByStudentId(ctx context.Context, studentId string) ([]*model.Submission, error)
	GetSubmissionById(ctx context.Context, submissionId string) (*model.Submission, error)
	EnrollCourse(ctx context.Context, enrollment *model.CourseEnrollment) error
	UpdateSubmission(ctx context.Context, submission *model.Submission) error
	List(ctx context.Context) ([]*model.Course, error)
}
