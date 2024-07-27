package location

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewLocationHandler(query LocationQuery, logError core.Log) *LocationHandler {
	paramIndex, filterIndex := search.BuildParams(reflect.TypeOf(LocationFilter{}))
	return &LocationHandler{query: query, logError: logError, paramIndex: paramIndex, filterIndex: filterIndex}
}

type LocationHandler struct {
	query       LocationQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *LocationHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		location, err := h.query.Load(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if location == nil {
			core.JSON(w, http.StatusNotFound, location)
		} else {
			core.JSON(w, http.StatusOK, location)
		}
	}
}

func (h *LocationHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := LocationFilter{Filter: &search.Filter{}}
	search.Decode(r, &filter, h.paramIndex, h.filterIndex)

	offset := search.GetOffset(filter.Limit, filter.Page)
	var users []Location
	users, total, err := h.query.Search(r.Context(), &filter, filter.Limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &search.Result{List: &users, Total: total})
}
