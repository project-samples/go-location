package tour

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"net/http"
	"reflect"
)

type TourHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewTourHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), load func(ctx context.Context, id interface{}, result interface{}) (bool, error), logError func(context.Context, string), writeLog func(context.Context, string, string, bool, string) error) TourHandler {
	searchModelType := reflect.TypeOf(TourFilter{})
	modelType := reflect.TypeOf(Tour{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &tourHandler{load: load, SearchHandler: searchHandler, Error: logError, Log: writeLog}
}

type tourHandler struct {
	load func(ctx context.Context, id interface{}, result interface{}) (bool, error)
	*search.SearchHandler
	Error func(context.Context, string)
	Log   func(context.Context, string, string, bool, string) error
}

func (h *tourHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		var tour Tour
		ok, err := h.load(r.Context(), id, &tour)
		sv.RespondIfFound(w, r, tour, ok, err, h.Error, nil)
	}
}
