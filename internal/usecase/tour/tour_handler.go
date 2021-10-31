package tour

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
)

type TourHandler struct {
	*sv.LoadHandler
	*search.SearchHandler
	Service TourService
}

func NewTourHandler(tourService TourService, logError func(context.Context, string)) *TourHandler {
	modelType := reflect.TypeOf(Tour{})
	searchModelType := reflect.TypeOf(TourFilter{})
	searchHandler := search.NewSearchHandler(tourService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewLoadHandler(tourService.Load, modelType, logError)
	return &TourHandler{LoadHandler: genericHandler, SearchHandler: searchHandler, Service: tourService}
}

func (h *TourHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
