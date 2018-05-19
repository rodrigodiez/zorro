package zorro

import (
	"sync"
)

type memory struct {
	mutex *sync.Mutex
	f     map[string]string
	b     map[string]string
}

func (m *memory) LoadOrStore(id string, mask string) (actualMask string, loaded bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	actual, ok := m.f[id]

	if ok {
		return actual, true
	}

	m.f[id], m.b[mask] = mask, id

	return mask, false
}

func (m *memory) Resolve(mask string) (id string, ok bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	id, ok = m.b[mask]

	return id, ok
}

// NewInMemoryStorage creates a new Storage that lives in memory.
//
// This type of storage is intented for testing and demonstration purposes only.
// Although the implementation is safe for concurrent use, it is not persisted.
func NewInMemoryStorage() Storage {
	return &memory{
		mutex: &sync.Mutex{},
		f:     make(map[string]string),
		b:     make(map[string]string),
	}
}
