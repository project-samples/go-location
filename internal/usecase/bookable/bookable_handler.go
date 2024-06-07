package bookable

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type BookableHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewBookableHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), service BookableService, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) BookableHandler {
	searchModelType := reflect.TypeOf(BookableFilter{})
	modelType := reflect.TypeOf(Bookable{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &bookableHandler{service: service, SearchHandler: searchHandler}
}

type bookableHandler struct {
	service BookableService
	*search.SearchHandler
}

func (h *bookableHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sv.JSON(w, sv.IsFound(res), res)
	}
}
