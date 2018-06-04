package memory

import (
	"testing"

	metricsMocks "github.com/rodrigodiez/zorro/lib/mocks/metrics"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestImplementsStorage(t *testing.T) {
	var _ storage.Storage = New()
}

func TestLoadOrStoreReturnsValueAndNilIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	mem := New()

	value, err := mem.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
	t.Parallel()

	mem := New()

	mem.LoadOrStore("foo", "bar")
	value, err := mem.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

func TestResolve(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name          string
		loadedKey     string
		loadedValue   string
		value         string
		expectedKey   string
		expectedError bool
	}{
		{name: "Key exists", loadedKey: "foo", loadedValue: "bar", value: "bar", expectedKey: "foo", expectedError: false},
		{name: "Key does not exist", loadedKey: "foo", loadedValue: "bar", value: "baz", expectedKey: "", expectedError: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mem := New()

			mem.LoadOrStore(tc.loadedKey, tc.loadedValue)
			key, err := mem.Resolve(tc.value)

			assert.Equal(t, tc.expectedKey, key)
			if tc.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestLoadOrStoreIncrementsStoreOpsCounterIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	storage := New().WithMetrics(&storage.Metrics{StoreOps: counter})
	storage.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func TestLoadOrStoreIncrementsLoadOpsCounterIfKeyExists(t *testing.T) {
	t.Parallel()
	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	storage := New().WithMetrics(&storage.Metrics{LoadOps: counter})
	storage.LoadOrStore("foo", "bar")
	storage.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}
func TestResolveIncrementsResolveOpsCounter(t *testing.T) {
	t.Parallel()
	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	storage := New().WithMetrics(&storage.Metrics{ResolveOps: counter})
	storage.Resolve("bar")

	counter.AssertCalled(t, "Add", int64(1))
}
