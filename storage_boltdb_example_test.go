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
	id, _ := storage.Resolve("bar")

	fmt.Printf("Id for 'bar' is '%s'\n", id)

	// Output:
	// Id for 'bar' is 'foo'
}

func ExampleNewBoltDBStorage_other() {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path) // We ignore the error this time. Naughty!
	defer storage.Close()

	storage.LoadOrStore("foo", "bar")

	mask, loaded := storage.LoadOrStore("foo", "42")
	id, _ := storage.Resolve("bar")

	fmt.Printf("Mask of 'foo' is '%s'\n", mask)
	fmt.Printf("Mask of 'foo' was loaded from storage: %t\n", loaded)
	fmt.Printf("Id for 'bar' is '%s'\n", id)

	_, ok := storage.Resolve("42")

	fmt.Printf("Id for '42' could be resolved: %t", ok)

	// Output:
	// Mask of 'foo' is 'bar'
	// Mask of 'foo' was loaded from storage: true
	// Id for 'bar' is 'foo'
	// Id for '42' could be resolved: false
}
