// Package zorro contains interfaces and concrete implementations
// to manage Zorro.
package zorro

// Zorro is the interface that wraps the methods to mask and unmask ids
type Zorro interface {
	Mask(id string) (mask string)
	Unmask(mask string) (id string, ok bool)
}

type zorro struct {
	generator Generator
	storage   Storage
}

func (t *zorro) Mask(id string) (mask string) {

	tmpMask := t.generator.Generate(id)

	mask, _ = t.storage.LoadOrStore(id, tmpMask)

	return mask
}

func (t *zorro) Unmask(mask string) (id string, ok bool) {

	return t.storage.Resolve(mask)
}

// New creates a new Zorro
func New(g Generator, s Storage) Zorro {
	return &zorro{
		generator: g,
		storage:   s,
	}
}
