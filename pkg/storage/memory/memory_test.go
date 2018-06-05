package memory_test

import (
	"testing"

	metricsMocks "github.com/rodrigodiez/zorro/lib/mocks/metrics"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestImplementsStorage(t *testing.T) {
	var _ storage.Storage = memory.New()
}

func TestLoadOrStoreReturnsValueAndNilIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	sto := memory.New()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
	t.Parallel()

	sto := memory.New()

	sto.LoadOrStore("foo", "bar")
	value, err := sto.LoadOrStore("foo", "baz")

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
			sto := memory.New()

			sto.LoadOrStore(tc.loadedKey, tc.loadedValue)
			key, err := sto.Resolve(tc.value)

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

	sto := memory.New().WithMetrics(&storage.Metrics{StoreOps: counter})
	sto.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func TestLoadOrStoreIncrementsLoadOpsCounterIfKeyExists(t *testing.T) {
	t.Parallel()
	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	sto := memory.New().WithMetrics(&storage.Metrics{LoadOps: counter})
	sto.LoadOrStore("foo", "bar")
	sto.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}
func TestResolveIncrementsResolveOpsCounter(t *testing.T) {
	t.Parallel()
	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	sto := memory.New().WithMetrics(&storage.Metrics{ResolveOps: counter})
	sto.Resolve("bar")

	counter.AssertCalled(t, "Add", int64(1))
}
