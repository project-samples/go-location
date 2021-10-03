package tour

import (
	"github.com/core-go/search"
	"github.com/core-go/service"
)

type TourService interface {
	search.SearchService
	service.ViewService
}
