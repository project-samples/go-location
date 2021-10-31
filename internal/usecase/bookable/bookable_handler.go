package bookable

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
)

type BookableHandler struct {
	*sv.LoadHandler
	*search.SearchHandler
	Service BookableService
}

func NewBookableHandler(bookableService BookableService, logError func(context.Context, string)) *BookableHandler {
	modelType := reflect.TypeOf(Bookable{})
	searchModelType := reflect.TypeOf(BookableFilter{})
	searchHandler := search.NewSearchHandler(bookableService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewLoadHandler(bookableService.Load, modelType, logError)
	return &BookableHandler{LoadHandler: genericHandler, SearchHandler: searchHandler, Service: bookableService}
}

func (h *BookableHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
