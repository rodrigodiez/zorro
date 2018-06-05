package memory

import (
	"expvar"
	"fmt"

	"github.com/rodrigodiez/zorro/pkg/storage"
)

func ExampleNew_other() {
	sto := New()

	sto.LoadOrStore("foo", "bar")

	value, _ := sto.LoadOrStore("foo", "42")
	key, _ := sto.Resolve("bar")

	fmt.Printf("Value of 'foo' is '%s'\n", value)
	fmt.Printf("Key for 'bar' is '%s'\n", key)

	// Output:
	// Value of 'foo' is 'bar'
	// Key for 'bar' is 'foo'
}

func ExampleNew() {
	sto := New()

	sto.LoadOrStore("foo", "bar")
	key, _ := sto.Resolve("bar")

	fmt.Printf("Key for 'bar' is '%s'\n", key)

	// Output:
	// Key for 'bar' is 'foo'
}

func ExampleNew_with_metrics() {
	loadOps := expvar.NewInt("loadOps")
	storeOps := expvar.NewInt("storeOps")
	resolveOps := expvar.NewInt("resolveOps")

	sto := New()
	sto.WithMetrics(&storage.Metrics{LoadOps: loadOps, StoreOps: storeOps, ResolveOps: resolveOps})

	sto.LoadOrStore("foo", "bar")
	fmt.Printf("Load: %d, Store: %d, Resolve: %d\n", loadOps.Value(), storeOps.Value(), resolveOps.Value())

	sto.LoadOrStore("foo", "bar")
	fmt.Printf("Load: %d, Store: %d, Resolve: %d\n", loadOps.Value(), storeOps.Value(), resolveOps.Value())

	sto.Resolve("bar")
	fmt.Printf("Load: %d, Store: %d, Resolve: %d\n", loadOps.Value(), storeOps.Value(), resolveOps.Value())

	// Output:
	// Load: 0, Store: 1, Resolve: 0
	// Load: 1, Store: 1, Resolve: 0
	// Load: 1, Store: 1, Resolve: 1
}
