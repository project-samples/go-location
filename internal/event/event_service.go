package event

import (
	"github.com/core-go/search"
	"github.com/core-go/service"
)

type EventService interface {
	search.SearchService
	service.ViewService
}
