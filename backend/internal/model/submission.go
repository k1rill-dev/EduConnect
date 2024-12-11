package model

import "time"

type Submission struct {
	Id          string    `bson:"_id"`
	Topic       string    `bson:"topic"`
	Assignment  string    `bson:"assignment"`
	CourseId    string    `bson:"course_id"`
	StudentId   string    `bson:"student_id"`
	Grade       string    `bson:"grade"`
	SubmittedAt time.Time `bson:"submission_at"`
}
