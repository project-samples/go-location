package tour

import "context"

type TourQuery interface {
	Load(ctx context.Context, id string) (*Tour, error)
	Search(ctx context.Context, filter *TourFilter, limit int64, offset int64) ([]Tour, int64, error)
}
