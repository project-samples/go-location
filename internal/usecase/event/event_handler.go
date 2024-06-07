package event

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type EventHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewEventHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), load func(ctx context.Context, id interface{}, result interface{}) (bool, error), logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) EventHandler {
	searchModelType := reflect.TypeOf(EventFilter{})
	modelType := reflect.TypeOf(Event{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &eventHandler{load: load, SearchHandler: searchHandler}
}

type eventHandler struct {
	load func(ctx context.Context, id interface{}, result interface{}) (bool, error)
	*search.SearchHandler
}

func (h *eventHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		var event Event
		ok, err := h.load(r.Context(), id, &event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if ok {
			sv.JSON(w, http.StatusOK, &event)
		} else {
			sv.JSON(w, http.StatusNotFound, nil)
		}
	}
}
