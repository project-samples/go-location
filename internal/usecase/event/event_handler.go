package event

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
)

type EventHandler struct {
	*sv.LoadHandler
	*search.SearchHandler
	Service EventService
}

func NewEventHandler(eventService EventService, logError func(context.Context, string)) *EventHandler {
	modelType := reflect.TypeOf(Event{})
	searchModelType := reflect.TypeOf(EventFilter{})
	searchHandler := search.NewSearchHandler(eventService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewLoadHandler(eventService.Load, modelType, logError)
	return &EventHandler{LoadHandler: genericHandler, SearchHandler: searchHandler, Service: eventService}
}

func (h *EventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
