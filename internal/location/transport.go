package location

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type LocationTranport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewLocationTransport(db *mongo.Database, logError func(context.Context, string, ...map[string]interface{})) LocationTranport {
	locationService := NewLocationQuery(db)
	locationHandler := NewLocationHandler(locationService, logError)
	return locationHandler
}
