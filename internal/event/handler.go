package event

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewEventHandler(query EventQuery, logError core.Log) *EventHandler {
	paramIndex, filterIndex := search.BuildParams(reflect.TypeOf(EventFilter{}))
	return &EventHandler{query: query, logError: logError, paramIndex: paramIndex, filterIndex: filterIndex}
}

type EventHandler struct {
	query       EventQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *EventHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		event, err := h.query.Load(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if event == nil {
			core.JSON(w, http.StatusNotFound, event)
		} else {
			core.JSON(w, http.StatusOK, event)
		}
	}
}

func (h *EventHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := EventFilter{Filter: &search.Filter{}}
	search.Decode(r, &filter, h.paramIndex, h.filterIndex)

	offset := search.GetOffset(filter.Limit, filter.Page)
	var users []Event
	users, total, err := h.query.Search(r.Context(), &filter, filter.Limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &search.Result{List: &users, Total: total})
}
