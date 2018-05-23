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
	counter := expvar.NewInt("zorro_mask_ops")

	zorro := New(uuid.NewV4(), memory.New())
	zorro.WithMetrics(&Metrics{MaskOps: counter})

	zorro.Mask("foo")
	fmt.Printf("Mask has been called %d times\n", counter.Value())
	zorro.Mask("bar")
	fmt.Printf("Mask has been called %d times\n", counter.Value())

	// Output:
	// Mask has been called 1 times
	// Mask has been called 2 times
}
