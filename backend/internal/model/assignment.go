package model

type Assignment struct {
	Title          string `bson:"title"`
	TheoryFile     string `bson:"theory_file"`
	AdditionalInfo string `bson:"additional_info"`
}
