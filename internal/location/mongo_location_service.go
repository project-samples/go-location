package location

import (
	"reflect"

	mgo "github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	"github.com/core-go/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoLocationService struct {
	search.SearchService
	service.ViewService
	Mapper mgo.Mapper
}

func NewLocationService(db *mongo.Database) *MongoLocationService {
	var model Location
	modelType := reflect.TypeOf(model)
	mapper := geo.NewMapper(modelType)
	queryBuilder := query.NewBuilder(modelType)
	searchService, genericService := mgo.NewSearchLoaderWithQuery(db, "location", modelType, queryBuilder.BuildQuery, search.GetSort, mapper.DbToModel)
	return &MongoLocationService{SearchService: searchService, ViewService: genericService, Mapper: mapper}
}
