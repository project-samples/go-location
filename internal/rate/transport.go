package rate

import (
	"context"
	"reflect"

	m "go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/mongo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo/query"
)

func NewRateTransport(db *m.Database, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) RateHandler {
	locationRateType := reflect.TypeOf(Rate{})
	locationRateQuery := query.UseQuery(locationRateType)

	locationRateSearchBuilder := mongo.NewSearchBuilder(db, "locationRate", locationRateQuery, search.GetSort)
	getLocationRate := mongo.UseGet(db, "locationRate", locationRateType)
	locationRateHandler := NewRateHandler(locationRateSearchBuilder.Search, getLocationRate, logError, writeLog)

	return locationRateHandler
}
