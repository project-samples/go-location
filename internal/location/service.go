package location

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mgo "github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	mq "github.com/core-go/search/mongo/query"
)

type LocationQuery interface {
	Load(ctx context.Context, id string) (*Location, error)
	Search(ctx context.Context, filter *LocationFilter, limit int64, offset int64) ([]Location, int64, error)
}

func NewLocationService(db *mongo.Database) LocationQuery {
	queryLocation := mq.UseQuery[Location, *LocationFilter]()
	mapper := geo.NewMapper[Location]()
	return &LocationService{db.Collection("location"), db.Collection("locationInfo"), mapper.DbToModel, queryLocation}
}

type LocationService struct {
	collection     *mongo.Collection
	infoCollection *mongo.Collection
	Map            func(location *Location)
	BuildQuery     func(filter *LocationFilter) (bson.D, bson.M)
}

func (s *LocationService) Load(ctx context.Context, id string) (*Location, error) {
	filter := bson.M{"_id": id}
	var location Location
	ok, err := mgo.FindOne(ctx, s.collection, filter, &location)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, err
	}
	s.Map(&location)
	var locationInfo LocationInfo
	ok, err = mgo.FindOne(ctx, s.infoCollection, filter, &locationInfo)
	if !ok {
		return &location, err
	} else {
		location.Info = &locationInfo
		return &location, err
	}
}
func (s *LocationService) Search(ctx context.Context, filter *LocationFilter, limit int64, offset int64) ([]Location, int64, error) {
	query, fields := s.BuildQuery(filter)
	var locations []Location
	total, err := s.collection.CountDocuments(ctx, query)
	if err != nil || total == 0 {
		return locations, total, err
	}
	opts := options.Find()
	if len(filter.Sort) > 0 {
		opts.SetSort(mgo.BuildSort(filter.Sort, reflect.TypeOf(LocationFilter{})))
	}
	opts.SetSkip(offset)
	if limit > 0 {
		opts.SetLimit(limit)
	}
	if fields != nil {
		opts.Projection = fields
	}
	cursor, err := s.collection.Find(ctx, query, opts)
	if err != nil {
		return locations, total, err
	}
	err = cursor.All(ctx, &locations)
	return locations, total, err
}
