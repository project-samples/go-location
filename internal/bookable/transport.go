package bookable

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/core"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
)

type BookableTranport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewBookableTransport(db *mongo.Database, logError core.Log) BookableTranport {
	bookableMapper := geo.NewMapper[Bookable]()
	queryBookable := mq.UseQuery[Bookable, *BookableFilter]()
	bookableQuery := query.NewQuery[Bookable, string, *BookableFilter](db, "bookable", queryBookable, search.GetSort, bookableMapper.DbToModel)
	bookableHandler := NewBookableHandler(bookableQuery, logError)
	return bookableHandler
}
