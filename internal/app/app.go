package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo"
	"github.com/teris-io/shortid"
	"reflect"

	"go-service/internal/usecase/bookable"
	"go-service/internal/usecase/event"
	"go-service/internal/usecase/location"
	"go-service/internal/usecase/tour"
)

type ApplicationContext struct {
	HealthHandler   *health.Handler
	LocationHandler location.LocationHandler
	EventHandler    event.EventHandler
	BookableHandler bookable.BookableHandler
	TourHandler     tour.TourHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	db, err := mongo.Setup(ctx, root.Mongo)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	mongoChecker := mongo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	locationType := reflect.TypeOf(location.Location{})
	locationMapper := geo.NewMapper(locationType)
	locationQuery := query.UseQuery(locationType)
	locationSearchBuilder := mongo.NewSearchBuilder(db, "location", locationQuery, search.GetSort, locationMapper.DbToModel)
	locationRepository := mongo.NewViewRepository(db, "location", locationType, locationMapper.DbToModel)
	locationService := location.NewLocationService(locationRepository)
	locationHandler := location.NewLocationHandler(locationSearchBuilder.Search, locationService, logError, nil)

	eventType := reflect.TypeOf(event.Event{})
	eventMapper := geo.NewMapper(eventType)
	eventQuery := query.UseQuery(eventType)
	eventSearchBuilder := mongo.NewSearchBuilder(db, "event", eventQuery, search.GetSort, eventMapper.DbToModel)
	eventRepository := mongo.NewViewRepository(db, "event", eventType, eventMapper.DbToModel)
	eventService := event.NewEventService(eventRepository)
	eventHandler := event.NewEventHandler(eventSearchBuilder.Search, eventService, logError, nil)

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
	tourRepository := mongo.NewViewRepository(db, "tour", tourType, tourMapper.DbToModel)
	tourService := tour.NewTourService(tourRepository)
	tourHandler := tour.NewTourHandler(tourSearchBuilder.Search, tourService, logError, nil)

	return &ApplicationContext{
		HealthHandler:   healthHandler,
		LocationHandler: locationHandler,
		EventHandler:    eventHandler,
		BookableHandler: bookableHandler,
		TourHandler:     tourHandler,
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
