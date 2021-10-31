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

func NewTourHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service TourService, logError func(context.Context, string), writeLog func(context.Context, string, string, bool, string) error) TourHandler {
	searchModelType := reflect.TypeOf(TourFilter{})
	modelType := reflect.TypeOf(Tour{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &tourHandler{service: service, SearchHandler: searchHandler, Error: logError, Log: writeLog}
}

type tourHandler struct {
	service TourService
	*search.SearchHandler
	Error func(context.Context, string)
	Log   func(context.Context, string, string, bool, string) error
}

func (h *tourHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
