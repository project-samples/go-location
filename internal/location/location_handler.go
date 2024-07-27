package location

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/core"
	"github.com/core-go/search"
)

type LocationHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewLocationHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), service LocationService, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) LocationHandler {
	searchModelType := reflect.TypeOf(LocationFilter{})
	modelType := reflect.TypeOf(Location{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &locationHandler{service: service, SearchHandler: searchHandler}
}

type locationHandler struct {
	service LocationService
	*search.SearchHandler
	Error func(context.Context, string)
	Log   func(context.Context, string, string, bool, string) error
}

func (h *locationHandler) Load(w http.ResponseWriter, r *http.Request) {
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
