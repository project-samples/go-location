package app

import (
	"context"
	"github.com/gorilla/mux"

	. "github.com/core-go/core/constants"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.HandleFunc("/health", app.Health.Check).Methods(GET)

	location := "/locations"
	r.HandleFunc(location, app.Location.Search).Methods(GET)
	r.HandleFunc(location+"/search", app.Location.Search).Methods(GET, POST)
	r.HandleFunc(location+"/{id}", app.Location.Load).Methods(GET)

	locationRate := "/location-rates"
	r.HandleFunc(locationRate+"/search", app.LocationRate.Search).Methods(GET, POST)
	r.HandleFunc(locationRate+"/{id}", app.LocationRate.Load).Methods(GET)

	event := "/events"
	r.HandleFunc(event+"/search", app.Event.Search).Methods(GET, POST)
	r.HandleFunc(event+"/{id}", app.Event.Load).Methods(GET)

	bookable := "/bookables"
	r.HandleFunc(bookable+"/search", app.Bookable.Search).Methods(GET, POST)
	r.HandleFunc(bookable+"/{id}", app.Bookable.Load).Methods(GET)

	tour := "/tours"
	r.HandleFunc(tour+"", app.Tour.Search).Methods(GET, POST)
	r.HandleFunc(tour+"/search", app.Tour.Search).Methods(GET, POST)
	r.HandleFunc(tour+"/{id}", app.Tour.Load).Methods(GET)

	return nil
}
