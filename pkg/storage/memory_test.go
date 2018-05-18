package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsStorage(t *testing.T) {
	var _ Storage = NewMem()
}

func TestLoadOrStoreReTestturnsMaskAndFalseIfIdDoesNotExist(t *testing.T) {
	mem := NewMem()

	mask, loaded := mem.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", mask)
	assert.Equal(t, false, loaded)
}

func TestLoadOrStoreReturnsActualMaskAndTrueIfKeyExists(t *testing.T) {
	mem := NewMem()

	mem.LoadOrStore("foo", "bar")
	mask, loaded := mem.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", mask)
	assert.Equal(t, true, loaded)
}

func TestResolve(t *testing.T) {
	tt := []struct {
		name       string
		loadedID   string
		loadedMask string
		mask       string
		expectedID string
		expectedOk bool
	}{
		{name: "Id exists", loadedID: "foo", loadedMask: "bar", mask: "bar", expectedID: "foo", expectedOk: true},
		{name: "Id does not exist", loadedID: "foo", loadedMask: "bar", mask: "baz", expectedID: "", expectedOk: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mem := NewMem()

			mem.LoadOrStore(tc.loadedID, tc.loadedMask)
			id, ok := mem.Resolve(tc.mask)

			assert.Equal(t, tc.expectedID, id)
			assert.Equal(t, tc.expectedOk, ok)
		})
	}
}
