package event

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"net/http"
	"reflect"
)

type EventHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewEventHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), load func(ctx context.Context, id interface{}, result interface{}) (bool, error), logError func(context.Context, string), writeLog func(context.Context, string, string, bool, string) error) EventHandler {
	searchModelType := reflect.TypeOf(EventFilter{})
	modelType := reflect.TypeOf(Event{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &eventHandler{load: load, SearchHandler: searchHandler, Error: logError, Log: writeLog}
}

type eventHandler struct {
	load func(ctx context.Context, id interface{}, result interface{}) (bool, error)
	*search.SearchHandler
	Error func(context.Context, string)
	Log   func(context.Context, string, string, bool, string) error
}

func (h *eventHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		var event Event
		ok, err := h.load(r.Context(), id, &event)
		if err == nil && !ok {
			sv.JSON(w, http.StatusNotFound, nil)
		} else {
			sv.RespondModel(w, r, event, err, h.Error, nil)
		}
	}
}
