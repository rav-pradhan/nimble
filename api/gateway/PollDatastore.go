package gateway

import "github.com/rav-pradhan/pollr/api/domain"

type PollDatastore interface {
	CreatePoll(request domain.CreatePollRequest) error
}
