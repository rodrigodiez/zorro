package zorro

import (
	"fmt"
	"log"
	"os"
)

func ExampleNewBoltDBStorage() {
	path := getTmpPath()
	defer os.Remove(path)

	storage, err := NewBoltDBStorage(path)

	if err != nil {
		log.Fatal(err)
	}

	defer storage.Close()

	storage.LoadOrStore("foo", "bar")
	key, _ := storage.Resolve("bar")

	fmt.Printf("Key for 'bar' is '%s'\n", key)

	// Output:
	// Key for 'bar' is 'foo'
}

func ExampleNewBoltDBStorage_other() {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path) // We ignore the error this time. Naughty!
	defer storage.Close()

	storage.LoadOrStore("foo", "bar")

	value, loaded := storage.LoadOrStore("foo", "42")
	key, _ := storage.Resolve("bar")

	fmt.Printf("Value of 'foo' is '%s'\n", value)
	fmt.Printf("Value of 'foo' was loaded from storage: %t\n", loaded)
	fmt.Printf("Key for 'bar' is '%s'\n", key)

	_, ok := storage.Resolve("42")

	fmt.Printf("Key for '42' could be resolved: %t", ok)

	// Output:
	// Value of 'foo' is 'bar'
	// Value of 'foo' was loaded from storage: true
	// Key for 'bar' is 'foo'
	// Key for '42' could be resolved: false
}
