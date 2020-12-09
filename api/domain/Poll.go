package domain

import (
	"time"
)

type Poll struct {
	Name                         string       `json:"name" bson:"name"`
	Description                  string       `json:"description" bson:"name"`
	MultipleOptionsCanBeSelected bool         `json:"multiple_options_can_be_selected" bson:"multiple_options_can_be_selected"`
	Options                      []string     `json:"options" bson:"options"`
	Respondents                  []Respondent `json:"respondents" bson:"respondents"`
	Creator                      string       `json:"creator" bson:"creator"`
	CreatedAt                    time.Time    `json:"created_at" bson:"created_at"`
	ClosesAt                     time.Time    `json:"closes_at" bson:"closes_at"`
}
