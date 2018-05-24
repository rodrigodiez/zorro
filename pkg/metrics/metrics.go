package metrics

// IntCounter is an interface metric counters backed up by an int64
type IntCounter interface {
	Add(int64)
}
