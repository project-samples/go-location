package app

import (
	"context"
	"reflect"

	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo/query"
	"github.com/teris-io/shortid"

	"go-service/internal/usecase/bookable"
	"go-service/internal/usecase/event"
	"go-service/internal/usecase/location"
	"go-service/internal/usecase/rate"
	"go-service/internal/usecase/tour"
)

type ApplicationContext struct {
	Health       *health.Handler
	Location     location.LocationHandler
	LocationRate rate.RateHandler
	Event        event.EventHandler
	Bookable     bookable.BookableHandler
	Tour         tour.TourHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := mongo.Setup(ctx, conf.Mongo)
	if err != nil {
		return nil, err
	}
	logError := log.LogError

	mongoChecker := mongo.NewHealthChecker(db.Client())
	healthHandler := health.NewHandler(mongoChecker)

	locationType := reflect.TypeOf(location.Location{})
	locationInfoType := reflect.TypeOf(location.LocationInfo{})

	locationMapper := geo.NewMapper(locationType)
	locationQuery := query.UseQuery(locationType)
	locationSearchBuilder := mongo.NewSearchBuilder(db, "location", locationQuery, search.GetSort, locationMapper.DbToModel)
	locationRepository := mongo.NewViewRepository(db, "location", locationType, locationMapper.DbToModel)
	locationInfoRepository := mongo.NewViewRepository(db, "locationInfo", locationInfoType)
	locationService := location.NewLocationService(db, locationMapper.DbToModel, locationRepository, locationInfoRepository)
	locationHandler := location.NewLocationHandler(locationSearchBuilder.Search, locationService, logError, nil)

	locationRateType := reflect.TypeOf(rate.Rate{})
	locationRateQuery := query.UseQuery(locationRateType)

	locationRateSearchBuilder := mongo.NewSearchBuilder(db, "locationRate", locationRateQuery, search.GetSort)
	getLocationRate := mongo.UseGet(db, "locationRate", locationRateType)
	locationRateHandler := rate.NewRateHandler(locationRateSearchBuilder.Search, getLocationRate, logError, nil)

	eventType := reflect.TypeOf(event.Event{})
	eventMapper := geo.NewMapper(eventType)
	eventQuery := query.UseQuery(eventType)
	eventSearchBuilder := mongo.NewSearchBuilder(db, "event", eventQuery, search.GetSort, eventMapper.DbToModel)
	getEvent := mongo.UseGet(db, "event", eventType, eventMapper.DbToModel)
	eventHandler := event.NewEventHandler(eventSearchBuilder.Search, getEvent, logError, nil)

	bookableType := reflect.TypeOf(bookable.Bookable{})
	bookableMapper := geo.NewMapper(bookableType)
	bookableQuery := query.UseQuery(bookableType)
	bookableSearchBuilder := mongo.NewSearchBuilder(db, "bookable", bookableQuery, search.GetSort, bookableMapper.DbToModel)
	bookableRepository := mongo.NewViewRepository(db, "bookable", bookableType, bookableMapper.DbToModel)
	bookableService := bookable.NewBookableService(bookableRepository)
	bookableHandler := bookable.NewBookableHandler(bookableSearchBuilder.Search, bookableService, logError, nil)

	tourType := reflect.TypeOf(tour.Tour{})
	tourMapper := geo.NewMapper(tourType)
	tourQuery := query.UseQuery(tourType)
	tourSearchBuilder := mongo.NewSearchBuilder(db, "tour", tourQuery, search.GetSort, tourMapper.DbToModel)
	getTour := mongo.UseGet(db, "tour", tourType, tourMapper.DbToModel)
	tourHandler := tour.NewTourHandler(tourSearchBuilder.Search, getTour, logError, nil)

	return &ApplicationContext{
		Health:       healthHandler,
		Location:     locationHandler,
		Event:        eventHandler,
		Bookable:     bookableHandler,
		Tour:         tourHandler,
		LocationRate: locationRateHandler,
	}, nil
}

var sid *shortid.Shortid

func ShortId() (string, error) {
	if sid == nil {
		s, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			return "", err
		}
		sid = s
	}
	return sid.Generate()
}
func Generate(ctx context.Context) (string, error) {
	return ShortId()
}
