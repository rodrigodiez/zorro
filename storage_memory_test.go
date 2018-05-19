package zorro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryImplementsStorage(t *testing.T) {
	var _ Storage = NewInMemoryStorage()
}

func TestLoadOrStoreReTestturnsValueAndFalseIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	mem := NewInMemoryStorage()

	value, loaded := mem.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Equal(t, false, loaded)
}

func TestMemoryLoadOrStoreReturnsActualValueAndTrueIfKeyExists(t *testing.T) {
	t.Parallel()

	mem := NewInMemoryStorage()

	mem.LoadOrStore("foo", "bar")
	value, loaded := mem.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", value)
	assert.Equal(t, true, loaded)
}

func TestMemoryResolve(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		loadedKey   string
		loadedValue string
		value       string
		expectedKey string
		expectedOk  bool
	}{
		{name: "Key exists", loadedKey: "foo", loadedValue: "bar", value: "bar", expectedKey: "foo", expectedOk: true},
		{name: "Key does not exist", loadedKey: "foo", loadedValue: "bar", value: "baz", expectedKey: "", expectedOk: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mem := NewInMemoryStorage()

			mem.LoadOrStore(tc.loadedKey, tc.loadedValue)
			key, ok := mem.Resolve(tc.value)

			assert.Equal(t, tc.expectedKey, key)
			assert.Equal(t, tc.expectedOk, ok)
		})
	}
}
