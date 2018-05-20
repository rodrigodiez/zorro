// Package service contains interfaces and concrete implementations
// to manage Zorro.
package service

import (
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

// Zorro is the interface that wraps the methods to mask and unmask keys
type Zorro interface {
	Mask(key string) (value string)
	Unmask(value string) (key string, ok bool)
}

type zorro struct {
	generator generator.Generator
	storage   storage.Storage
}

func (t *zorro) Mask(key string) (value string) {

	tmpValue := t.generator.Generate(key)

	value, _ = t.storage.LoadOrStore(key, tmpValue)

	return value
}

func (t *zorro) Unmask(value string) (key string, ok bool) {

	return t.storage.Resolve(value)
}

// New creates a new Zorro
func New(g generator.Generator, s storage.Storage) Zorro {
	return &zorro{
		generator: g,
		storage:   s,
	}
}
