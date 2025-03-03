package caches

import (
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type (
	Interface[T any] interface {
		// Remember value by key. If key is exists and ttl is not expired then return value without call callback.
		// If key is not exists or ttl is expired then call callback and remember value by key.
		// Error is returned from callback if failed.
		Remember(key string, ttlS uint64, callback func() (T, error)) (T, error)

		// Clear all cache.
		Clear()

		// Forget cache value by key.
		Forget(key string) (T, error)

		// Determine your own logic with function to forget cache by key
		ForgetFn(func(key string) (T, error)) (T, error)

		// Get length of cache.
		Length() uint64

		// Get cache value by key
		Get(key string) (T, error)

		// Get keys that stored on caches
		Keys() []string
	}

	item[T any] struct {
		TTL  uint64
		Data T
	}

	StorageType string

	Config struct {
		StorageType StorageType
		Dir         string
	}
)

const (
	StorageTypeMemory = StorageType("memory")
)

func Init[T any](cfg Config) (Interface[T], error) {
	switch cfg.StorageType {
	case StorageTypeMemory:
		return InitMemory[T](cfg), nil
	default:
		return nil, errors.NewWithCode(codes.CodeNotImplemented, "storage type '%s' not implemented!", cfg.StorageType)
	}
}
