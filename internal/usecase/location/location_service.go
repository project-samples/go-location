package location

import (
	"context"

	sv "github.com/core-go/service"
)

type LocationService interface {
	Load(ctx context.Context, id string) (*Location, error)
}

func NewLocationService(repository sv.ViewRepository) LocationService {
	return &locationService{repository: repository}
}

type locationService struct {
	repository sv.ViewRepository
}

func (s *locationService) Load(ctx context.Context, id string) (*Location, error) {
	var location Location
	ok, err := s.repository.LoadAndDecode(ctx, id, &location)
	if !ok {
		return nil, err
	} else {
		return &location, err
	}
}
