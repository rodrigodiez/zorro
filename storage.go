package zorro

// Storage is the interface that wraps the methods to load, store and resolve
// ids and masks.
type Storage interface {
	LoadOrStore(id string, mask string) (actualMask string, loaded bool)
	Resolve(mask string) (id string, ok bool)
}

// StorageCloser is an interface for Storage that allows underlying resources
// to be closed
type StorageCloser interface {
	Storage
	Close()
}
