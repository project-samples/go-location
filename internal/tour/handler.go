package tour

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewTourHandler(query TourQuery, logError core.Log) *TourHandler {
	paramIndex, filterIndex := search.BuildAttributes(reflect.TypeOf(TourFilter{}))
	return &TourHandler{query: query, logError: logError, paramIndex: paramIndex, filterIndex: filterIndex}
}

type TourHandler struct {
	query       TourQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *TourHandler) Load(w http.ResponseWriter, r *http.Request) {
	id, err := core.GetRequiredString(w, r)
	if err == nil {
		tour, err := h.query.Load(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if tour == nil {
			core.JSON(w, http.StatusNotFound, tour)
		} else {
			core.JSON(w, http.StatusOK, tour)
		}
	}
}

func (h *TourHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := TourFilter{Filter: &search.Filter{}}
	err := search.Decode(r, &filter, h.paramIndex, h.filterIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset := search.GetOffset(filter.Limit, filter.Page)
	var users []Tour
	users, total, err := h.query.Search(r.Context(), &filter, filter.Limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &search.Result{List: &users, Total: total})
}
