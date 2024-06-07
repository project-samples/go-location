package location

import (
	"context"
	sv "github.com/core-go/core"
)

type LocationService interface {
	Load(ctx context.Context, id string) (*Location, error)
}

func NewLocationService(repository sv.ViewRepository, repositoryInfo sv.ViewRepository) LocationService {
	return &locationService{repository: repository, repositoryInfo: repositoryInfo}
}

type locationService struct {
	repository     sv.ViewRepository
	repositoryInfo sv.ViewRepository
}

func (s *locationService) Load(ctx context.Context, id string) (*Location, error) {
	var location Location
	ok, err := s.repository.Get(ctx, id, &location)
	if !ok {
		return nil, err
	}
	var locationInfo LocationInfo
	ok, err = s.repositoryInfo.Get(ctx, id, &locationInfo)
	if !ok {
		return &location, err
	} else {
		location.Info = &locationInfo
		return &location, err
	}
}
