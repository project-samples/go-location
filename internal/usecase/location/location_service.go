package location

import (
	"context"
	sv "github.com/core-go/core"
	q "github.com/core-go/mongo/query"
	b "github.com/core-go/mongo/query/builder"

	"github.com/core-go/search"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationService interface {
	Load(ctx context.Context, id string) (*Location, error)
	Search(ctx context.Context, filter *LocationFilter, limit int64, skip int64) ([]Location, int64, error)
}

func NewLocationService(db *mongo.Database, mp func(context.Context, interface{}) (interface{}, error), repository sv.ViewRepository, repositoryInfo sv.ViewRepository) LocationService {
	buildQuery := b.UseQuery[Location, *LocationFilter]()
	se := q.NewSearchBuilder[Location, *LocationFilter](db, "location", buildQuery, search.GetSort, mp)
	return &locationService{SearchBuilder: se, repository: repository, repositoryInfo: repositoryInfo}
}

type locationService struct {
	*q.SearchBuilder[Location, *LocationFilter]
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
