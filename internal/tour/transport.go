package tour

import (
	"context"
	"reflect"

	m "go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo/query"
)

func NewTourTransport(db *m.Database, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) TourHandler {
	tourType := reflect.TypeOf(Tour{})
	tourMapper := geo.NewMapper(tourType)
	tourQuery := query.UseQuery(tourType)
	tourSearchBuilder := mongo.NewSearchBuilder(db, "tour", tourQuery, search.GetSort, tourMapper.DbToModel)
	getTour := mongo.UseGet(db, "tour", tourType, tourMapper.DbToModel)
	tourHandler := NewTourHandler(tourSearchBuilder.Search, getTour, logError, writeLog)
	return tourHandler
}
