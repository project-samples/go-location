package bookable

import (
	"context"

	sv "github.com/core-go/service"
)

type BookableService interface {
	Load(ctx context.Context, id string) (*Bookable, error)
}

func NewBookableService(repository sv.ViewRepository) BookableService {
	return &bookableService{repository: repository}
}

type bookableService struct {
	repository sv.ViewRepository
}

func (s *bookableService) Load(ctx context.Context, id string) (*Bookable, error) {
	var Bookable Bookable
	ok, err := s.repository.LoadAndDecode(ctx, id, &Bookable)
	if !ok {
		return nil, err
	} else {
		return &Bookable, err
	}
}
