package rate

import (
	"context"
	sv "github.com/core-go/service"
)

type RateService interface {
	Load(ctx context.Context, id string) (*Rate, error)
}

func NewLocationRateService(repository sv.ViewRepository) RateService {
	return &rateService{repository: repository}
}

type rateService struct {
	repository sv.ViewRepository
}

func (s *rateService) Load(ctx context.Context, id string) (*Rate, error) {
	var locationRate Rate
	ok, err := s.repository.LoadAndDecode(ctx, id, &locationRate)
	if !ok {
		return nil, err
	} else {
		return &locationRate, err
	}
}
