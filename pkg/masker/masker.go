// Package masker contains interfaces and concrete implementations
// to manage masking and unmasking of ids and masks.
package masker

import (
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

// Masker is the interface that wraps the methods to mask and unmask ids
type Masker interface {
	Mask(id string) (mask string)
	Unmask(mask string) (id string, ok bool)
}

type masker struct {
	generator generator.Generator
	storage   storage.Storage
}

func (t *masker) Mask(id string) (mask string) {

	tmpMask := t.generator.Generate(id)
	mask, _ = t.storage.LoadOrStore(id, tmpMask)

	return mask
}

func (t *masker) Unmask(mask string) (id string, ok bool) {

	return t.storage.Resolve(mask)
}

// New creates a new Masker. A Storage and Generator must be provided
func New(g generator.Generator, s storage.Storage) Masker {
	return &masker{
		generator: g,
		storage:   s,
	}
}
