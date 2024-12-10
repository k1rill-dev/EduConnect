package model

import "time"

type Assignment struct {
	Id          string    `bson:"_id"`
	CourseId    string    `bson:"course_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
}
