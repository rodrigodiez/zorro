package zorro

import (
	"sync"
)

type memory struct {
	mutex *sync.RWMutex
	k     map[string]string
	v     map[string]string
}

func (m *memory) LoadOrStore(key string, value string) (actualValue string, loaded bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	actual, ok := m.k[key]

	if ok {
		return actual, true
	}

	m.k[key], m.v[value] = value, key

	return value, false
}

func (m *memory) Resolve(value string) (key string, ok bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	key, ok = m.v[value]

	return key, ok
}

// NewInMemoryStorage creates a new Storage that lives in memory.
//
// This type of storage is intended for testing and demonstration purposes only.
// Although the implementation is safe for concurrent use, it is not persisted.
func NewInMemoryStorage() Storage {
	return &memory{
		mutex: &sync.RWMutex{},
		k:     make(map[string]string),
		v:     make(map[string]string),
	}
}
