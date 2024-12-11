package requests

import (
	"EduConnect/internal/model"
	"time"
)

type SignUpRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Picture   string `json:"picture" validate:"required"`
	Bio       string `json:"bio" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokensRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignOutRequest struct {
}

type CreateCourseRequest struct {
	Title       string         `json:"title" form:"title" validate:"required"`
	Description string         `json:"description" form:"description" validate:"required"`
	TeacherId   string         `json:"teacher_id" form:"teacher_id" validate:"required"`
	StartDate   time.Time      `json:"start_date" form:"start_date" validate:"required"`
	EndDate     time.Time      `json:"end_date" form:"end_date" validate:"required"`
	Topics      []TopicRequest `json:"topics" form:"topics"`
}

type GetCourseByIdRequest struct {
	Id string `json:"_id" validate:"required"`
}

type TopicRequest struct {
	Title       string              `json:"title" form:"title" validate:"required"`
	Assignments []AssignmentRequest `json:"assignments" form:"assignments"`
}

type AssignmentRequest struct {
	Title          string `json:"title" form:"title" validate:"required"`
	TheoryFile     string `json:"theory_file" form:"theory_file" validate:"required"`
	AdditionalInfo string `json:"additional_info" form:"additional_info" validate:"required"`
}

type UpdateRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `bson:"first_name" validate:"required"`
	Surname   string `bson:"surname" validate:"required"`
	Picture   string `json:"picture" validate:"required"`
	Bio       string `json:"bio" validate:"required"`
}

type CreateJobRequest struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Location    string `bson:"location"`
}

type UpdateJobRequest struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Location    string `bson:"location"`
}

type CreateJobApplication struct {
	CompanyId string `bson:"company_id" validate:"required"`
	StudentId string `bson:"student_id" validate:"required"`
	Status    string `bson:"status"  validate:"required"`
}

type CreatePortfolioRequest struct {
	StudentId string                 `bson:"student_id"`
	Items     []model.PortfolioItems `bson:"items"`
}

type SubmitAssignmentRequest struct {
	Topic      string `json:"topic" form:"topic" validate:"required"`
	Assignment string `json:"assignment" form:"assignment" validate:"required"`
	CourseId   string `json:"course_id" form:"course_id" validate:"required"`
	Submission string `json:"submission" form:"course_id" validate:"required"`
	// TheoryFile string `json:"theory_file" form:"theory_file" validate:"required"`
	// Grade      string `bson:"grade" validate:"required"`
	// SubmittedAt time.Time `bson:"submission_at"`
	// StudentId   string    `bson:"student_id" `
}
