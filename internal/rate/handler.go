package rate

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewRateHandler(query RateQuery, logError core.Log) *RateHandler {
	paramIndex, filterIndex := search.BuildParams(reflect.TypeOf(RateFilter{}))
	return &RateHandler{query: query, logError: logError, paramIndex: paramIndex, filterIndex: filterIndex}
}

type RateHandler struct {
	query       RateQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *RateHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
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
	search.Decode(r, &filter, h.paramIndex, h.filterIndex)

	offset := search.GetOffset(filter.Limit, filter.Page)
	var users []Rate
	users, total, err := h.query.Search(r.Context(), &filter, filter.Limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &search.Result{List: &users, Total: total})
}
