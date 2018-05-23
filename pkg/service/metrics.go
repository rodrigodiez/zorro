package service

// IntCounter is an interface metric counters backed up by an int64
type IntCounter interface {
	Add(int64)
}

// Metrics contains references to user provided metrics
//
// MaskOps: Number of times Mask() has been called
// UnmaskOps: Number of times Unmask() has been called
type Metrics struct {
	MaskOps   IntCounter
	UnmaskOps IntCounter
}

// WithMetrics allows user to configure Zorro to emit operational metrics
func (z *zorro) WithMetrics(m *Metrics) Zorro {
	z.metrics = m

	return z
}
