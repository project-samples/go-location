package location

import (
	"github.com/core-go/search"
	"github.com/core-go/service"
)

type LocationService interface {
	search.SearchService
	service.ViewService
}
