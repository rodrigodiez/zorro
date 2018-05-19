// Package zorro contains interfaces and concrete implementations
// to manage Zorro.
package zorro

// Zorro is the interface that wraps the methods to mask and unmask keys
type Zorro interface {
	Mask(key string) (value string)
	Unmask(value string) (key string, ok bool)
}

type zorro struct {
	generator Generator
	storage   Storage
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
func New(g Generator, s Storage) Zorro {
	return &zorro{
		generator: g,
		storage:   s,
	}
}
