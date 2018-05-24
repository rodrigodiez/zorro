package storage

import "github.com/rodrigodiez/zorro/pkg/metrics"

// Storage is the interface that wraps the methods to load, store and resolve
// keys and values.
type Storage interface {
	LoadOrStore(key string, value string) (string, bool)
	Resolve(value string) (string, bool)
	WithMetrics(metrics *Metrics) Storage
}

// Closer is an interface to free up underlying resources
type Closer interface {
	Storage
	Close()
}

// Metrics contains references to user provided metrics
//
// LoadOps: Number of times value was loaded for key
// StoreOps: Number of times value was stored for key
type Metrics struct {
	LoadOps    metrics.IntCounter
	StoreOps   metrics.IntCounter
	ResolveOps metrics.IntCounter
}
