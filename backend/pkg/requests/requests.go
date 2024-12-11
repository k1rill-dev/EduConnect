package requests

import "time"

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Picture  string `json:"picture" validate:"required"`
	Bio      string `json:"bio" validate:"required"`
	Role     string `json:"role" validate:"required"`
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

type TopicRequest struct {
	Title       string              `json:"title" form:"title" validate:"required"`
	Assignments []AssignmentRequest `json:"assignments" form:"assignments"`
}

type AssignmentRequest struct {
	Title          string `json:"title" form:"title" validate:"required"`
	TheoryFile     string `json:"theory_file" form:"theory_file" validate:"required"`
	AdditionalInfo string `json:"additional_info" form:"additional_info" validate:"required"`
}
