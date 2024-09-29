package bookable

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewBookableHandler(query BookableQuery, logError core.Log) *BookableHandler {
	paramIndex, filterIndex := search.BuildAttributes(reflect.TypeOf(BookableFilter{}))
	return &BookableHandler{query: query, logError: logError, paramIndex: paramIndex, filterIndex: filterIndex}
}

type BookableHandler struct {
	query       BookableQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *BookableHandler) Load(w http.ResponseWriter, r *http.Request) {
	id, err := core.GetRequiredString(w, r)
	if err == nil {
		bookable, err := h.query.Load(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if bookable == nil {
			core.JSON(w, http.StatusNotFound, bookable)
		} else {
			core.JSON(w, http.StatusOK, bookable)
		}
	}
}

func (h *BookableHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := BookableFilter{Filter: &search.Filter{}}
	err := search.Decode(r, &filter, h.paramIndex, h.filterIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset := search.GetOffset(filter.Limit, filter.Page)
	var users []Bookable
	users, total, err := h.query.Search(r.Context(), &filter, filter.Limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &search.Result{List: &users, Total: total})
}
