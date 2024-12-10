package model

import "time"

type User struct {
	Id        string    `bson:"_id"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Picture   string    `bson:"picture"`
	CreatedAt time.Time `bson:"created_at"`
	Role      string    `bson:"role"`
}

func NewUser(id string, email string, password string, createdAt time.Time, role string) *User {
	return &User{Id: id, Role: role}
}
