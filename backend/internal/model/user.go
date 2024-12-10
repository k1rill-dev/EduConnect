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
	Bio       string           `bson:"bio"`
	CreatedAt time.Time        `bson:"created_at"`
	Role      string           `bson:"role"`
}

func NewUser(id string, email *values.Email, password *values.Password, picture string, bio string, createdAt time.Time, role string) *User {
	return &User{Id: id, Email: email, Password: password, Picture: picture, Bio: bio, CreatedAt: createdAt, Role: role}
}

// func (u *User) ToUserDb() *UserDB {
// 	return &UserDB{
// 		Id:        u.Id,
// 		Email:     u.Email.Email,
// 		Password:  u.Password.Password,
// 		Picture:   u.Picture,
// 		CreatedAt: u.CreatedAt,
// 		Role:      u.Role,
// 	}
// }
