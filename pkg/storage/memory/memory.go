package memory

import (
	"sync"

	"github.com/rodrigodiez/zorro/pkg/storage"
)

type memory struct {
	mutex   *sync.RWMutex
	k       map[string]string
	v       map[string]string
	metrics *storage.Metrics
}

func (m *memory) LoadOrStore(key string, value string) (actualValue string, loaded bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	actual, ok := m.k[key]

	if ok {
		m.incrLoadOps()
		return actual, true
	}

	m.k[key], m.v[value] = value, key
	m.incrStoreOps()

	return value, false
}

func (m *memory) Resolve(value string) (key string, ok bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	key, ok = m.v[value]
	m.incrResolveOps()

	return key, ok
}

// New creates a new Storage that lives in memory.
//
// This type of storage is intended for testing and demonstration purposes only.
// Although the implementation is safe for concurrent use, it is not persisted.
func New() storage.Storage {
	return &memory{
		mutex: &sync.RWMutex{},
		k:     make(map[string]string),
		v:     make(map[string]string),
	}
}

func (m *memory) WithMetrics(metrics *storage.Metrics) storage.Storage {
	m.metrics = metrics

	return m
}

func (m *memory) Close() {
}

func (m *memory) incrStoreOps() {
	if m.metrics != nil && m.metrics.StoreOps != nil {
		m.metrics.StoreOps.Add(int64(1))
	}
}

func (m *memory) incrLoadOps() {
	if m.metrics != nil && m.metrics.LoadOps != nil {
		m.metrics.LoadOps.Add(int64(1))
	}
}

func (m *memory) incrResolveOps() {
	if m.metrics != nil && m.metrics.ResolveOps != nil {
		m.metrics.ResolveOps.Add(int64(1))
	}
}
