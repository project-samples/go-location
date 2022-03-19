package location

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"net/http"
	"reflect"
)

type LocationHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewLocationHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), load func(ctx context.Context, id interface{}, result interface{}) (bool, error), logError func(context.Context, string), writeLog func(context.Context, string, string, bool, string) error) LocationHandler {
	searchModelType := reflect.TypeOf(LocationFilter{})
	modelType := reflect.TypeOf(Location{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &locationHandler{load: load, SearchHandler: searchHandler, Error: logError, Log: writeLog}
}

type locationHandler struct {
	load func(ctx context.Context, id interface{}, result interface{}) (bool, error)
	*search.SearchHandler
	Error func(context.Context, string)
	Log   func(context.Context, string, string, bool, string) error
}

func (h *locationHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		var location Location
		ok, err := h.load(r.Context(), id, &location)
		if err == nil && !ok {
			sv.JSON(w, http.StatusNotFound, nil)
		} else {
			sv.RespondModel(w, r, location, err, h.Error, nil)
		}
	}
}
