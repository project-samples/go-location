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
	var Tour Tour
	ok, err := s.repository.LoadAndDecode(ctx, id, &Tour)
	if !ok {
		return nil, err
	} else {
		return &Tour, err
	}
}
