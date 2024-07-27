package event

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/core"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
)

type EventTranport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewEventTransport(db *mongo.Database, logError core.Log) EventTranport {
	eventMapper := geo.NewMapper[Event]()
	queryEvent := mq.UseQuery[Event, *EventFilter]()
	eventQuery := query.NewQuery[Event, string, *EventFilter](db, "event", queryEvent, search.GetSort, eventMapper.DbToModel)
	eventHandler := NewEventHandler(eventQuery, logError)
	return eventHandler
}
