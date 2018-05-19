package zorro

import (
	"fmt"
)

func ExampleNewInMemoryStorage_other() {
	mem := NewInMemoryStorage()

	mem.LoadOrStore("foo", "bar")

	mask, loaded := mem.LoadOrStore("foo", "42")
	id, _ := mem.Resolve("bar")

	fmt.Printf("Mask of 'foo' is '%s'\n", mask)
	fmt.Printf("Mask of 'foo' was loaded from storage: %t\n", loaded)
	fmt.Printf("Id for 'bar' is '%s'\n", id)

	_, ok := mem.Resolve("42")

	fmt.Printf("Id for '42' could be resolved: %t", ok)

	// Output:
	// Mask of 'foo' is 'bar'
	// Mask of 'foo' was loaded from storage: true
	// Id for 'bar' is 'foo'
	// Id for '42' could be resolved: false
}

func ExampleNewInMemoryStorage() {
	mem := NewInMemoryStorage()

	mem.LoadOrStore("foo", "bar")
	id, _ := mem.Resolve("bar")

	fmt.Printf("Id for 'bar' is '%s'\n", id)

	// Output:
	// Id for 'bar' is 'foo'
}
