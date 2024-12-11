package requests

import "EduConnect/internal/model"

type SignUpRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `bson:"first_name" validate:"required"`
	Surname   string `bson:"surname" validate:"required"`
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

type SignOutRequest struct {
}
