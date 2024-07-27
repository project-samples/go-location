package rate

import "context"

type RateQuery interface {
	Load(ctx context.Context, id string) (*Rate, error)
	Search(ctx context.Context, filter *RateFilter, limit int64, offset int64) ([]Rate, int64, error)
}
