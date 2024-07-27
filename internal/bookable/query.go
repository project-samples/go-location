package bookable

import "context"

type BookableQuery interface {
	Load(ctx context.Context, id string) (*Bookable, error)
	Search(ctx context.Context, filter *BookableFilter, limit int64, offset int64) ([]Bookable, int64, error)
}
