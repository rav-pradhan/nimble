package domain

type CreatePollRequest struct {
	Name                    string   `json:"name" bson:"name" validate:"nonzero"`
	Description             string   `json:"description" bson:"name"`
	AllowMultipleSelections bool     `json:"allow_multiple_selections"`
	Options                 []string `json:"options" bson:"options" validate:"min=2"`
	Creator                 string   `json:"creator" bson:"creator" validate:"nonzero"`
}
