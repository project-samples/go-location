package bookable

import (
	"reflect"

	mgo "github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	"github.com/core-go/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookableService interface {
	search.SearchService
	service.ViewService
}

type MongoBookableService struct {
	search.SearchService
	service.ViewService
	Mapper mgo.Mapper
}

func NewBookableService(db *mongo.Database) *MongoBookableService {
	var model Bookable
	modelType := reflect.TypeOf(model)
	mapper := geo.NewMapper(modelType)
	queryBuilder := query.NewBuilder(modelType)
	searchService, viewService := mgo.NewSearchLoaderWithQuery(db, "bookable", modelType, queryBuilder.BuildQuery, search.GetSort, mapper.DbToModel)
	return &MongoBookableService{SearchService: searchService, ViewService: viewService, Mapper: mapper}
}
