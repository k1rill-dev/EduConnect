package model

import "time"

type Submission struct {
	Id           string    `bson:"_id"`
	AssignmentId string    `bson:"assignment_id"`
	StudentId    string    `bson:"student_id"`
	Grade        string    `bson:"grade"`
	SubmittedAt  time.Time `bson:"submission_at"`
}
