package safemap

import (
	"sync"
)

// SafeMap is a thread-safe map implementation.

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// NewSafeMap creates a new instance of SafeMap.

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		mu:   sync.RWMutex{},
		data: make(map[K]V),
	}
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.data[key]

	if !ok {
		return val, KeyError{
			Key: key,
		}
	}

	return val, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.data[key]

	if !ok {
		return KeyError{
			Key: key,
		}
	}

	m.data[key] = value

	return nil
}

// HasKey checks if the map contains the given key.

func (m *SafeMap[K, V]) HasKey(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.data[key]

	return ok

}

// KeyError is the error type for key-related errors.

type KeyError struct {
	Key interface{}
}

func (e KeyError) Error() string {
	return "Key error: " + e.Key.(string)
}
