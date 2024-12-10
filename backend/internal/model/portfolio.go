package model

type Portfolio struct {
	Id        string `bson:"_id"`
	StudentId string `bson:"student_id"`
}
