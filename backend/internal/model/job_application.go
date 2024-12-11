package model

import "time"

type JobApplication struct {
	Id        string    `bson:"_id,omitempty"`
	CompanyId string    `bson:"company_id"`
	StudentId string    `bson:"student_id"`
	Status    string    `bson:"status"`
	createdAt time.Time `bson:"created_at"`
}

func NewJobApplication(id string, companyId string, studentId string, status string) *JobApplication {
	return &JobApplication{
		Id:        id,
		CompanyId: companyId,
		StudentId: studentId,
		Status:    status,
		createdAt: time.Now(),
	}
}
