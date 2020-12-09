package domain

type CreateRespondentRequest struct {
	Name string `json:"name" bson:"name"`
}
