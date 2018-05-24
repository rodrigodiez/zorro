package service

import (
	"expvar"
	"fmt"

	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func ExampleNew() {
	zorro := New(uuid.NewV4(), memory.New())
	zorro.Mask("foo")
}

func ExampleNew_with_metrics() {
	maskOps := expvar.NewInt("maskOps")
	unmaskOps := expvar.NewInt("unmaskOps")

	zorro := New(uuid.NewV4(), memory.New())
	zorro.WithMetrics(&Metrics{MaskOps: maskOps, UnmaskOps: unmaskOps})

	zorro.Mask("foo")
	fmt.Printf("Mask: %d, Unmask: %d\n", maskOps.Value(), unmaskOps.Value())

	zorro.Unmask("bar")
	fmt.Printf("Mask: %d, Unmask: %d\n", maskOps.Value(), unmaskOps.Value())

	// Output:
	// Mask: 1, Unmask: 0
	// Mask: 1, Unmask: 1
}
