package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/mongo"
	"github.com/teris-io/shortid"

	"go-service/internal/usecase/bookable"
	"go-service/internal/usecase/event"
	"go-service/internal/usecase/location"
	"go-service/internal/usecase/tour"
)

type ApplicationContext struct {
	HealthHandler   *health.Handler
	LocationHandler *location.LocationHandler
	EventHandler    *event.EventHandler
	BookableHandler *bookable.BookableHandler
	TourHandler     *tour.TourHandler
}

func NewApp(ctx context.Context, mongoConfig mongo.MongoConfig) (*ApplicationContext, error) {
	db, err := mongo.Setup(ctx, mongoConfig)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	mongoChecker := mongo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	locationService := location.NewLocationService(db)
	locationHandler := location.NewLocationHandler(locationService, logError)
	eventService := event.NewEventService(db)
	eventHandler := event.NewEventHandler(eventService, logError)
	bookableService := bookable.NewBookableService(db)
	bookableHandler := bookable.NewBookableHandler(bookableService, logError)
	tourService := tour.NewTourService(db)
	tourHandler := tour.NewTourHandler(tourService, logError)
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
