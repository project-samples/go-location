package event

import (
	"context"
	"reflect"

	m "go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo/query"
)

func NewEventTransport(db *m.Database, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) EventHandler {
	eventType := reflect.TypeOf(Event{})
	eventMapper := geo.NewMapper(eventType)
	eventQuery := query.UseQuery(eventType)
	eventSearchBuilder := mongo.NewSearchBuilder(db, "event", eventQuery, search.GetSort, eventMapper.DbToModel)
	getEvent := mongo.UseGet(db, "event", eventType, eventMapper.DbToModel)
	eventHandler := NewEventHandler(eventSearchBuilder.Search, getEvent, logError, writeLog)
	return eventHandler
}
