package tour

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type TourHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
}

func NewTourHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), load func(ctx context.Context, id interface{}, result interface{}) (bool, error), logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) TourHandler {
	searchModelType := reflect.TypeOf(TourFilter{})
	modelType := reflect.TypeOf(Tour{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &tourHandler{load: load, SearchHandler: searchHandler}
}

type tourHandler struct {
	load func(ctx context.Context, id interface{}, result interface{}) (bool, error)
	*search.SearchHandler
}

func (h *tourHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		var res Tour
		ok, err := h.load(r.Context(), id, &res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if ok {
			sv.JSON(w, http.StatusOK, &res)
		} else {
			sv.JSON(w, http.StatusNotFound, nil)
		}
	}
}
