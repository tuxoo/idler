package cache

import (
	"context"
	"errors"
	"sync"
)

type MemoryCache[K comparable, V any] struct {
	cache map[K]V
	sync.RWMutex
}

func NewMemoryCache[K comparable, V any]() *MemoryCache[K, V] {
	return &MemoryCache[K, V]{
		cache: make(map[K]V),
	}
}

func (c *MemoryCache[K, V]) Set(ctx context.Context, key K, value V) {
	c.Lock()
	c.cache[key] = value
	c.Unlock()
}

func (c *MemoryCache[K, V]) Get(ctx context.Context, key K) (V, error) {
	c.RLock()
	value, exist := c.cache[key]
	c.RUnlock()

	if !exist {
		return value, errors.New("not found")
	}

	return value, nil
}
