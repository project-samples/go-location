package event

import "context"

type EventQuery interface {
	Load(ctx context.Context, id string) (*Event, error)
	Search(ctx context.Context, filter *EventFilter, limit int64, offset int64) ([]Event, int64, error)
}
