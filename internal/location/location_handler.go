package location

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
)

type LocationHandler struct {
	*sv.LoadHandler
	*search.SearchHandler
	Service LocationService
}

func NewLocationHandler(locationService LocationService, logError func(context.Context, string)) *LocationHandler {
	modelType := reflect.TypeOf(Location{})
	searchModelType := reflect.TypeOf(LocationSM{})
	searchHandler := search.NewSearchHandler(locationService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewLoadHandler(locationService.Load, modelType, logError)
	return &LocationHandler{LoadHandler: genericHandler, SearchHandler: searchHandler, Service: locationService}
}

func (h *LocationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
