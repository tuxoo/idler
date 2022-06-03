package cache

import "context"

type Cache[K comparable, V any] interface {
	Set(ctx context.Context, key K, value *V)
	Get(ctx context.Context, key K) (*V, error)
}
