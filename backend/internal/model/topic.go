package model

type Topic struct {
	Title       string        `bson:"title"`
	Assignments []*Assignment `bson:"assignments"`
}
