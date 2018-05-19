package zorro

import (
	"fmt"
)

func ExampleNewInMemoryStorage_other() {
	mem := NewInMemoryStorage()

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

func ExampleNewInMemoryStorage() {
	mem := NewInMemoryStorage()

	mem.LoadOrStore("foo", "bar")
	key, _ := mem.Resolve("bar")

	fmt.Printf("Key for 'bar' is '%s'\n", key)

	// Output:
	// Key for 'bar' is 'foo'
}
