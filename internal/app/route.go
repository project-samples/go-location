package app

import (
	"context"

	. "github.com/core-go/service"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	locationPath := "/locations"
	location := app.LocationHandler
	// r.HandleFunc(locationPath, location.GetAll).Methods(GET)
	r.HandleFunc(locationPath+"/search", location.Search).Methods(GET, POST)
	r.HandleFunc(locationPath+"/{id}", location.Load).Methods(GET)

	locationRatePath := "/locationsrate"
	locationRate := app.LocationRateHandler
	// r.HandleFunc(locationPath, location.GetAll).Methods(GET)
	r.HandleFunc(locationRatePath+"/search", locationRate.Search).Methods(GET, POST)
	r.HandleFunc(locationRatePath+"/{id}", locationRate.Load).Methods(GET)

	eventPath := "/events"
	event := app.EventHandler
	// r.HandleFunc(eventPath, event.GetAll).Methods(GET)
	r.HandleFunc(eventPath+"/search", event.Search).Methods(GET, POST)
	r.HandleFunc(eventPath+"/{id}", event.Load).Methods(GET)

	bookablePath := "/bookables"
	bookable := app.BookableHandler
	// r.HandleFunc(bookablePath, bookable.GetAll).Methods(GET)
	r.HandleFunc(bookablePath+"/search", bookable.Search).Methods(GET, POST)
	r.HandleFunc(bookablePath+"/{id}", bookable.Load).Methods(GET)

	tourPath := "/tours"
	tour := app.TourHandler
	// r.HandleFunc(tourPath, tour.GetAll).Methods(GET)
	r.HandleFunc(tourPath+"", tour.Search).Methods(GET, POST)
	r.HandleFunc(tourPath+"/search", tour.Search).Methods(GET, POST)
	r.HandleFunc(tourPath+"/{id}", tour.Load).Methods(GET)

	return nil
}
