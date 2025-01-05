package caches

import (
	"time"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type (
	Interface[T any] interface {
		// Remember value by key. If key is exists and ttl is not expired then return value without call callback.
		// If key is not exists or ttl is expired then call callback and remember value by key.
		Remember(key string, ttlS uint64, callback func() T) (T, error)

		// Clear all cache.
		Clear()

		// Forget cache value by key.
		Forget(key string) (T, error)

		// Get length of cache.
		Length() uint64
	}

	cache[T any] struct {
		Dir     string
		Storage map[string]item[T]
	}

	item[T any] struct {
		TTL  uint64
		Data T
	}
)

func NewCache[T any]() Interface[T] {
	return &cache[T]{
		Dir:     "",
		Storage: map[string]item[T]{},
	}
}

func (c *cache[T]) Remember(key string, ttlS uint64, callback func() T) (T, error) {
	if i, isOk := c.Storage[key]; isOk {
		if i.TTL > uint64(time.Now().Unix()) {
			return i.Data, nil
		} else {
			delete(c.Storage, key)
			return c.Remember(key, ttlS, callback)
		}
	} else {
		data := callback()
		c.Storage[key] = item[T]{
			TTL:  uint64(time.Now().Unix()) + ttlS,
			Data: data,
		}
		return data, nil
	}
}

func (c *cache[T]) Clear() {
	c.Storage = map[string]item[T]{}
}

func (c *cache[T]) Forget(key string) (T, error) {
	if item, isOk := c.Storage[key]; isOk {
		delete(c.Storage, key)
		return item.Data, nil
	} else {
		return item.Data, errors.NewWithCode(codes.CodeNotFound, "key not found")
	}
}

func (c *cache[T]) Length() uint64 {
	return uint64(len(c.Storage))
}
