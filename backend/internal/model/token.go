package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccountClaims struct {
	Jti       string `json:"jti"`
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
	Sub       string `json:"sub"`
	Email     string `json:"email"`
	DeviceId  string `json:"device_id"`
	Iss       string `json:"iss"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type Token struct {
	Id        string `bson:"_id, omitempty"`
	DeviceId  string `bson:"deviceId"`
	AccountId string `bson:"accountId"`
	Token     string `bson:"token"`
	ExpiresAt int64  `bson:"expiresAt"`
	// TokenType string        `bson:"tokenType"`
}

func NewToken(id, deviceId, accountId, token string, expiresAt int64) *Token {
	return &Token{Id: id, DeviceId: deviceId, AccountId: accountId, Token: token, ExpiresAt: expiresAt}
}
