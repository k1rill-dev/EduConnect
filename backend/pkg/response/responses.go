package response

type RefreshTokensResponse struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignUpResponse struct {
	Id           string `json:"id" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	FirstName    string `json:"first_name" validate:"required"`
	Surname      string `json:"surname" validate:"required"`
	Picture      string `json:"picture" validate:"required"`
	Bio          string `json:"bio" validate:"required"`
	Role         string `json:"role" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignInResponse struct {
	Id           string `json:"id" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	FirstName    string `json:"first_name" validate:"required"`
	Surname      string `json:"surname" validate:"required"`
	Picture      string `json:"picture" validate:"required"`
	Bio          string `json:"bio" validate:"required"`
	Role         string `json:"role" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignOutResponse struct {
}

type UpdateResponse struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `bson:"first_name" validate:"required"`
	Surname   string `bson:"surname" validate:"required"`
	Picture   string `json:"picture" validate:"required"`
	Bio       string `json:"bio" validate:"required"`
}

type ErrorResponse struct {
	Error string `json:"error" validate:"required"`
}

type SuccessResponse struct {
	Message string `json:"message" validate:"required"`
}
