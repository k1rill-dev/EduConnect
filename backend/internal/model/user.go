package model

import (
	"EduConnect/internal/values"
	"time"
)

type User struct {
	Id        string           `bson:"_id"`
	Email     *values.Email    `bson:"email"`
	Password  *values.Password `bson:"password"`
	Picture   string           `bson:"picture"`
	CreatedAt time.Time        `bson:"created_at"`
	Role      string           `bson:"role"`
}

func NewUser(id string, email *values.Email, password *values.Password, picture string, createdAt time.Time, role string) *User {
	return &User{Id: id, Email: email, Password: password, Picture: picture, CreatedAt: createdAt, Role: role}
}
