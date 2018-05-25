package metrics

// IntCounter is an interface metric counters backed up by an int64
type IntCounter interface {
	Add(int64)
}

// Service contains references to user provided metrics
//
// MaskOps: Number of times Mask() has been called
// UnmaskOps: Number of times Unmask() has been called
type ServiceMetrics struct {
	MaskOps   IntCounter
	UnmaskOps IntCounter
}
