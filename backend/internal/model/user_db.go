package model

// import (
// 	"EduConnect/internal/values"
// 	"time"
// )

// type UserDB struct {
// 	Id        string    `bson:"_id"`
// 	Email     string    `bson:"email"`
// 	Password  []byte    `bson:"password"`
// 	Picture   string    `bson:"picture"`
// 	CreatedAt time.Time `bson:"created_at"`
// 	Role      string    `bson:"role"`
// }

// func (u *UserDB) ToUser() *User {
// 	email := &values.Email{u.Email}
// 	password := &values.Password{u.Password}
// 	return &User{
// 		Id:        u.Id,
// 		Email:     email,
// 		Password:  password,
// 		Picture:   u.Picture,
// 		CreatedAt: u.CreatedAt,
// 		Role:      u.Role,
// 	}
// }
