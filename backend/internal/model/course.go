package model

import "time"

type Course struct {
	Id          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	TeacherId   string    `bson:"teacher_id"`
	StartDate   time.Time `bson:"start_date"`
	EndDate     time.Time `bson:"end_date"`
	Topics      []*Topic  `bson:"topics"`
	CreatedAt   time.Time `bson:"created_at"`
}
