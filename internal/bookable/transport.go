package bookable

import (
	"context"
	"reflect"

	m "go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo/query"
)

func NewBookableTransport(db *m.Database, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) BookableHandler {
	bookableType := reflect.TypeOf(Bookable{})
	bookableMapper := geo.NewMapper(bookableType)
	bookableQuery := query.UseQuery(bookableType)
	bookableSearchBuilder := mongo.NewSearchBuilder(db, "bookable", bookableQuery, search.GetSort, bookableMapper.DbToModel)
	bookableRepository := mongo.NewViewRepository(db, "bookable", bookableType, bookableMapper.DbToModel)
	bookableService := NewBookableService(bookableRepository)
	bookableHandler := NewBookableHandler(bookableSearchBuilder.Search, bookableService, logError, writeLog)
	return bookableHandler
}
