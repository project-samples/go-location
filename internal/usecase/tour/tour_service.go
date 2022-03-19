package tour

import (
	"context"

	sv "github.com/core-go/service"
)

type TourService interface {
	Load(ctx context.Context, id string) (*Tour, error)
}

func NewTourService(repository sv.ViewRepository) TourService {
	return &tourService{repository: repository}
}

type tourService struct {
	repository sv.ViewRepository
}

func (s *tourService) Load(ctx context.Context, id string) (*Tour, error) {
	var tour Tour
	ok, err := s.repository.LoadAndDecode(ctx, id, &tour)
	if !ok {
		return nil, err
	} else {
		return &tour, err
	}
}
