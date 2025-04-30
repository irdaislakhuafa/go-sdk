package caches

import (
	"time"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type (
	memory[T any] struct {
		cfg     Config
		storage map[string]item[T]
	}
)

func InitMemory[T any](cfg Config) Interface[T] {
	return &memory[T]{
		cfg:     cfg,
		storage: map[string]item[T]{},
	}
}

func (m *memory[T]) Remember(key string, ttlS uint64, callback func() (T, error)) (T, error) {
	if i, isOk := m.storage[key]; isOk {
		if i.TTL > uint64(time.Now().Unix()) {
			return i.Data, nil
		} else {
			delete(m.storage, key)
			return m.Remember(key, ttlS, callback)
		}
	} else {
		data, err := callback()
		if err != nil {
			return data, err
		}

		m.storage[key] = item[T]{
			TTL:  uint64(time.Now().Unix()) + ttlS,
			Data: data,
		}
		return data, nil
	}
}

func (m *memory[T]) Clear() {
	m.storage = map[string]item[T]{}
}

func (m *memory[T]) Forget(key string) (T, error) {
	if item, isOk := m.storage[key]; isOk {
		delete(m.storage, key)
		return item.Data, nil
	} else {
		return item.Data, errors.NewWithCode(codes.CodeCacheKeyNotFound, "key not found")
	}
}

func (m *memory[T]) Length() uint64 {
	return uint64(len(m.storage))
}

func (m *memory[T]) ForgetFn(fn func(key string) (T, error)) (T, error) {
	var res T
	var err error
	for k := range m.storage {
		if res, err = fn(k); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (m *memory[T]) Get(key string) (T, error) {
	value, isExist := m.storage[key]
	if !isExist {
		return value.Data, errors.NewWithCode(codes.CodeCacheKeyNotFound, "cache with key '%s' not found", key)
	}

	return value.Data, nil
}

func (m *memory[T]) Keys() []string {
	keys := []string{}
	for k := range m.storage {
		keys = append(keys, k)
	}
	return keys
}
