package model

import "time"

type Course struct {
	Id          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	TeacherId   string    `bson:"teacher_id"`
	CreatedAt   time.Time `bson:"created_at"`
}
