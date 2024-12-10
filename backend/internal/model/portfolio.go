package model

type PortfolioItems struct {
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	URL         *string `bson:"url,omitempty"`
}

type Portfolio struct {
	Id        string           `bson:"_id"`
	StudentId string           `bson:"student_id"`
	Items     []PortfolioItems `bson:"items"`
}

func NewPortfolio(id string, studentId string, items []PortfolioItems) *Portfolio {
	return &Portfolio{
		Id:        id,
		StudentId: studentId,
		Items:     items,
	}
}
