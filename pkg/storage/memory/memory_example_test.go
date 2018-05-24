package memory

import (
	"expvar"
	"fmt"

	"github.com/rodrigodiez/zorro/pkg/storage"
)

func ExampleNew_other() {
	mem := New()

	mem.LoadOrStore("foo", "bar")

	value, loaded := mem.LoadOrStore("foo", "42")
	key, _ := mem.Resolve("bar")

	fmt.Printf("Value of 'foo' is '%s'\n", value)
	fmt.Printf("Value of 'foo' was loaded from storage: %t\n", loaded)
	fmt.Printf("Key for 'bar' is '%s'\n", key)

	_, ok := mem.Resolve("42")

	fmt.Printf("Key for '42' could be resolved: %t", ok)

	// Output:
	// Value of 'foo' is 'bar'
	// Value of 'foo' was loaded from storage: true
	// Key for 'bar' is 'foo'
	// Key for '42' could be resolved: false
}

func ExampleNew() {
	mem := New()

	mem.LoadOrStore("foo", "bar")
	key, _ := mem.Resolve("bar")

	fmt.Printf("Key for 'bar' is '%s'\n", key)

	// Output:
	// Key for 'bar' is 'foo'
}

func ExampleNew_with_metrics() {
	loadOps := expvar.NewInt("loadOps")
	storeOps := expvar.NewInt("storeOps")
	resolveOps := expvar.NewInt("resolveOps")

	memory := New()
	memory.WithMetrics(&storage.Metrics{LoadOps: loadOps, StoreOps: storeOps, ResolveOps: resolveOps})

	memory.LoadOrStore("foo", "bar")
	fmt.Printf("Load: %d, Store: %d, Resolve: %d\n", loadOps.Value(), storeOps.Value(), resolveOps.Value())

	memory.LoadOrStore("foo", "bar")
	fmt.Printf("Load: %d, Store: %d, Resolve: %d\n", loadOps.Value(), storeOps.Value(), resolveOps.Value())

	memory.Resolve("bar")
	fmt.Printf("Load: %d, Store: %d, Resolve: %d\n", loadOps.Value(), storeOps.Value(), resolveOps.Value())

	// Output:
	// Load: 0, Store: 1, Resolve: 0
	// Load: 1, Store: 1, Resolve: 0
	// Load: 1, Store: 1, Resolve: 1
}
