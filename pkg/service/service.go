// Package service contains interfaces and concrete implementations
// to manage Zorro.
package service

import (
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/metrics"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

// Metrics contains references to user provided metrics
//
// MaskOps: Number of times Mask() has been called
// UnmaskOps: Number of times Unmask() has been called
type Metrics struct {
	MaskOps   metrics.IntCounter
	UnmaskOps metrics.IntCounter
}

// Middleware is an interface to wrap calls to Zorro service
type Middleware func(Zorro) Zorro

// Zorro is the interface that wraps the methods to mask and unmask keys
type Zorro interface {
	Mask(key string) (value string)
	Unmask(value string) (key string, ok bool)
	WithMetrics(*Metrics) Zorro
}

type zorro struct {
	generator generator.Generator
	storage   storage.Storage
	metrics   *Metrics
}

func (z *zorro) Mask(key string) (value string) {
	z.incrMaskOps()

	tmpValue := z.generator.Generate(key)

	value, _ = z.storage.LoadOrStore(key, tmpValue)

	return value
}

func (z *zorro) Unmask(value string) (key string, ok bool) {

	z.incrUnmaskOps()

	return z.storage.Resolve(value)
}

func (z *zorro) incrMaskOps() {
	if z.metrics != nil && z.metrics.MaskOps != nil {
		z.metrics.MaskOps.Add(int64(1))
	}
}

func (z *zorro) incrUnmaskOps() {
	if z.metrics != nil && z.metrics.UnmaskOps != nil {
		z.metrics.UnmaskOps.Add(int64(1))
	}
}

// New creates a new Zorro
func New(g generator.Generator, s storage.Storage) Zorro {
	return &zorro{
		generator: g,
		storage:   s,
	}
}

// WithMetrics allows user to configure Zorro to emit operational metrics
func (z *zorro) WithMetrics(m *Metrics) Zorro {
	z.metrics = m

	return z
}
