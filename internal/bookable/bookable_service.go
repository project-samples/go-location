package bookable

import (
	"github.com/core-go/search"
	"github.com/core-go/service"
)

type BookableService interface {
	search.SearchService
	service.ViewService
}
