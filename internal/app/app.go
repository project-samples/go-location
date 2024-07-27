package app

import (
	"context"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/core-go/health"
	hm "github.com/core-go/health/mongo"
	"github.com/core-go/log"
	"github.com/teris-io/shortid"

	"go-service/internal/bookable"
	"go-service/internal/event"
	"go-service/internal/location"
	"go-service/internal/rate"
	"go-service/internal/tour"
)

type ApplicationContext struct {
	Health       *health.Handler
	Location     location.LocationTranport
	LocationRate rate.RateTranport
	Event        event.EventTranport
	Bookable     bookable.BookableTranport
	Tour         tour.TourTranport
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	client, err := m.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.Mongo.Database)
	logError := log.LogError

	mongoChecker := hm.NewHealthChecker(client)
	healthHandler := health.NewHandler(mongoChecker)

	locationHandler := location.NewLocationTransport(db, logError)
	locationRateHandler := rate.NewRateTransport(db, logError)
	eventHandler := event.NewEventTransport(db, logError)
	bookableHandler := bookable.NewBookableTransport(db, logError)
	tourHandler := tour.NewTourTransport(db, logError)

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
