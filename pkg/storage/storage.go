package storage

// Storage is the interface that wraps the methods to load, store and resolve
// keys and values.
type Storage interface {
	LoadOrStore(key string, value string) (actualValue string, loaded bool)
	Resolve(value string) (key string, ok bool)
}

// Closer is an interface to free up underlying resources
type Closer interface {
	Storage
	Close()
}
