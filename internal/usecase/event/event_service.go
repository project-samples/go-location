package event

import (
	"context"

	sv "github.com/core-go/service"
)

type EventService interface {
	Load(ctx context.Context, id string) (*Event, error)
}

func NewEventService(repository sv.ViewRepository) EventService {
	return &eventService{repository: repository}
}

type eventService struct {
	repository sv.ViewRepository
}

func (s *eventService) Load(ctx context.Context, id string) (*Event, error) {
	var event Event
	ok, err := s.repository.LoadAndDecode(ctx, id, &event)
	if !ok {
		return nil, err
	} else {
		return &event, err
	}
}
