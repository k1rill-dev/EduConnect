package repository

import (
	"EduConnect/internal/model"
	"context"
	// "nftvc-auth/internal/model"
)

type JwtRepository interface {
	SaveAccessToken(ctx context.Context, token *model.Token) error
	SaveRefreshToken(ctx context.Context, token *model.Token) error
	RevokeTokens(ctx context.Context, accountId string, deviceId string, acceptedToken string) error
	IsRevokedToken(ctx context.Context, accountId string, deviceId string, accessToken string) bool
	DeleteRefreshToken(ctx context.Context, accountId string, deviceId string) error
	GetAccessToken(ctx context.Context, accountId string, deviceId string) (string, error)
	CheckExistRefresh(ctx context.Context, refreshToken string) bool
	GetRefreshToken(ctx context.Context, accountId string, deviceId string) (string, error)
}
