package model

import "time"

type CourseEnrollment struct {
	Id         string    `bson:"_id"`
	CourseId   string    `bson:"course_id"`
	StudentId  string    `bson:"student_id"`
	EnrolledAt time.Time `bson:"enrolled_at"`
}
