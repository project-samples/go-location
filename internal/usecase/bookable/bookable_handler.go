package bookable

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"net/http"
	"reflect"
)

type BookableHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewBookableHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service BookableService, logError func(context.Context, string), writeLog func(context.Context, string, string, bool, string) error) BookableHandler {
	searchModelType := reflect.TypeOf(BookableFilter{})
	modelType := reflect.TypeOf(Bookable{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &bookableHandler{service: service, SearchHandler: searchHandler, Error: logError, Log: writeLog}
}

type bookableHandler struct {
	service BookableService
	*search.SearchHandler
	Error func(context.Context, string)
	Log   func(context.Context, string, string, bool, string) error
}

func (h *bookableHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
