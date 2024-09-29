package rate

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewRateHandler(query RateQuery, logError core.Log) *RateHandler {
	paramIndex, filterIndex := search.BuildAttributes(reflect.TypeOf(RateFilter{}))
	return &RateHandler{query: query, logError: logError, paramIndex: paramIndex, filterIndex: filterIndex}
}

type RateHandler struct {
	query       RateQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *RateHandler) Load(w http.ResponseWriter, r *http.Request) {
	id, err := core.GetRequiredString(w, r)
	if err == nil {
		rate, err := h.query.Load(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if rate == nil {
			core.JSON(w, http.StatusNotFound, rate)
		} else {
			core.JSON(w, http.StatusOK, rate)
		}
	}
}

func (h *RateHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := RateFilter{Filter: &search.Filter{}}
	err := search.Decode(r, &filter, h.paramIndex, h.filterIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset := search.GetOffset(filter.Limit, filter.Page)
	var users []Rate
	users, total, err := h.query.Search(r.Context(), &filter, filter.Limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &search.Result{List: &users, Total: total})
}
