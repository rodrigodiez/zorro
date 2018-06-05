package memory

import (
	"errors"
	"sync"

	"github.com/rodrigodiez/zorro/pkg/storage"
)

type memory struct {
	mutex   *sync.RWMutex
	k       map[string]string
	v       map[string]string
	metrics *storage.Metrics
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

func (sto *memory) LoadOrStore(key string, value string) (string, error) {
	sto.mutex.Lock()
	defer sto.mutex.Unlock()

	actual, ok := sto.k[key]

	if ok {
		sto.incrLoadOps()
		return actual, nil
	}

	sto.k[key], sto.v[value] = value, key
	sto.incrStoreOps()

	return value, nil
}

func (m *memory) Resolve(value string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	key, ok := m.v[value]
	m.incrResolveOps()

	if !ok {
		return "", errors.New("Key does not exist")
	}

	return key, nil
}

func (sto *memory) WithMetrics(metrics *storage.Metrics) storage.Storage {
	sto.metrics = metrics

	return sto
}

// Close is noop
func (sto *memory) Close() {
}

func (sto *memory) incrStoreOps() {
	if sto.metrics != nil && sto.metrics.StoreOps != nil {
		sto.metrics.StoreOps.Add(int64(1))
	}
}

func (sto *memory) incrLoadOps() {
	if sto.metrics != nil && sto.metrics.LoadOps != nil {
		sto.metrics.LoadOps.Add(int64(1))
	}
}

func (sto *memory) incrResolveOps() {
	if sto.metrics != nil && sto.metrics.ResolveOps != nil {
		sto.metrics.ResolveOps.Add(int64(1))
	}
}
