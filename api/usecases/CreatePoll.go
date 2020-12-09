package usecases

import (
	"github.com/rav-pradhan/nimble/api/domain"
	"github.com/rav-pradhan/nimble/api/gateway"
	"gopkg.in/validator.v2"
)

func CreatePoll(request domain.CreatePollRequest, datastore gateway.PollDatastore) error {
	err := validator.Validate(request)
	if err != nil {
		return err
	}
	return datastore.CreatePoll(request)
}
