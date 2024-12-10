package requests

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
