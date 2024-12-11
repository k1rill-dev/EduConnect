package model

import "time"

type Job struct {
	Id          string    `bson:"_id"`
	EmployerId  string    `bson:"employer_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Location    string    `bson:"location"`
	CreatedAt   time.Time `bson:"created_at"`
}

func NewJob(id string, employerId string, title string, description string, location string) *Job {
	return &Job{
		Id:          id,
		EmployerId:  employerId,
		Title:       title,
		Description: description,
		Location:    location,
		CreatedAt:   time.Now(),
	}
}
