package gateway

import "github.com/rav-pradhan/nimble/api/domain"

type PollDatastore interface {
	CreatePoll(request domain.CreatePollRequest) error
}
