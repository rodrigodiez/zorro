package zorro

// Storage is the interface that wraps the methods to load, store and resolve
// keys and values.
type Storage interface {
	LoadOrStore(key string, value string) (actualValue string, loaded bool)
	Resolve(value string) (key string, ok bool)
}

// StorageCloser is an interface for Storage that allows shut down of underlying resources
type StorageCloser interface {
	Storage
	Close()
}
