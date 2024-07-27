package location

import (
	"context"
	"reflect"

	m "go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo/query"
)

func NewLocationTransport(db *m.Database, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) LocationHandler {
	locationType := reflect.TypeOf(Location{})
	locationMapper := geo.NewMapper(locationType)
	locationQuery := query.UseQuery(locationType)
	locationSearchBuilder := mongo.NewSearchBuilder(db, "location", locationQuery, search.GetSort, locationMapper.DbToModel)
	locationRepository := mongo.NewViewRepository(db, "location", locationType, locationMapper.DbToModel)
	locationInfoRepository := mongo.NewViewRepository(db, "locationInfo", reflect.TypeOf(LocationInfo{}))
	locationService := NewLocationService(db, locationMapper.DbToModel, locationRepository, locationInfoRepository)
	locationHandler := NewLocationHandler(locationSearchBuilder.Search, locationService, logError, writeLog)
	return locationHandler
}
