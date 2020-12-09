package usecases_test

import (
	"testing"

	"github.com/rav-pradhan/nimble/api/domain"
	"github.com/rav-pradhan/nimble/api/usecases"
)

func TestCreatePoll(t *testing.T) {
	t.Run("An invalid CreatePollRequest returns an error at the validation stage", func(t *testing.T) {
		request := domain.CreatePollRequest{
			Name:                    "",
			Description:             "This is optional.",
			AllowMultipleSelections: false,
			Options:                 []string{},
			Creator:                 "",
		}
		result := usecases.CreatePoll(request, nil)
		if result == nil {
			t.Errorf("Expected error in validating request, got nil")
		}
	})
	t.Run("A valid CreatePollRequest means the datastore CreatePoll method is called", func(t *testing.T) {
		request := domain.CreatePollRequest{
			Name:                    "What to watch",
			Description:             "Let's work out what to watch next",
			AllowMultipleSelections: true,
			Options: []string{
				"A Game of Thrones",
				"The Wire",
				"Mad Men",
			},
			Creator: "Don Draper",
		}

		mockStore := NewMockPollDatastore()
		err := usecases.CreatePoll(request, &mockStore)
		if err != nil {
			t.Errorf("Expected nil, got validation error")
		}

		if !mockStore.CreatePollWasCalled {
			t.Errorf("The datastore's CreatePoll method was not called")
		}
	})
}

type MockPollDatastore struct {
	CreatePollWasCalled bool
}

func NewMockPollDatastore() MockPollDatastore {
	return MockPollDatastore{
		CreatePollWasCalled: false,
	}
}

func (m *MockPollDatastore) CreatePoll(request domain.CreatePollRequest) error {
	m.CreatePollWasCalled = true
	return nil
}
