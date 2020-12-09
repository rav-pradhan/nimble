package domain

type Respondent struct {
	Name    string   `json:"name" bson:"name"`
	Choices []string `json:"choices" bson:"choices"`
}
