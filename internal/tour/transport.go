package tour

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/core"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
)

type TourTranport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewTourTransport(db *mongo.Database, logError core.Log) TourTranport {
	queryTour := mq.UseQuery[Tour, *TourFilter]()
	tourQuery := query.NewQuery[Tour, string, *TourFilter](db, "tour", queryTour, search.GetSort)
	tourHandler := NewTourHandler(tourQuery, logError)
	return tourHandler
}
