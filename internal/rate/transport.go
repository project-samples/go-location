package rate

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/core"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
)

type RateTranport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewRateTransport(db *mongo.Database, logError core.Log) RateTranport {
	queryRate := mq.UseQuery[Rate, *RateFilter]()
	rateQuery := query.NewQuery[Rate, string, *RateFilter](db, "rate", queryRate, search.GetSort)
	rateHandler := NewRateHandler(rateQuery, logError)
	return rateHandler
}
