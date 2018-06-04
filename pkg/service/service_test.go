package service

import (
	"errors"
	"testing"

	generatorMocks "github.com/rodrigodiez/zorro/lib/mocks/generator"
	metricsMocks "github.com/rodrigodiez/zorro/lib/mocks/metrics"
	storageMocks "github.com/rodrigodiez/zorro/lib/mocks/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewReturnsZorro(t *testing.T) {
	t.Parallel()

	var _ Zorro = New(&generatorMocks.Generator{}, &storageMocks.Storage{})
}

func TestMask(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		key    string
		value  string
		actual string
		err    error
	}{
		{name: "value was not loaded by storage", key: "foo", value: "bar", actual: "bar", err: nil},
		{name: "value was loaded by storage", key: "foo", value: "bar", actual: "baz", err: nil},
		{name: "storage returned error", key: "foo", value: "bar", actual: "", err: errors.New("")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			generator := &generatorMocks.Generator{}
			storage := &storageMocks.Storage{}

			zorro := New(generator, storage)

			generator.On("Generate", tc.key).Return(tc.value).Once()
			storage.On("LoadOrStore", tc.key, tc.value).Return(tc.actual, tc.err).Once()

			value, err := zorro.Mask(tc.key)

			assert.Equal(t, tc.actual, value)
			assert.Equal(t, tc.err, err)

			storage.AssertExpectations(t)
		})
	}
}

func TestUnmask(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name  string
		key   string
		value string
		err   error
	}{
		{name: "value was unmasked", value: "bar", key: "foo", err: nil},
		{name: "value was not found", value: "bar", key: "", err: errors.New("")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			storage := &storageMocks.Storage{}

			zorro := New(&generatorMocks.Generator{}, storage)

			storage.On("Resolve", tc.value).Return(tc.key, tc.err).Once()

			key, err := zorro.Unmask(tc.value)

			assert.Equal(t, tc.key, key)
			assert.Equal(t, tc.err, err)

			storage.AssertExpectations(t)
		})
	}
}

func TestMaskIncrementsMaskOpCounter(t *testing.T) {
	generator := &generatorMocks.Generator{}
	storage := &storageMocks.Storage{}
	counter := &metricsMocks.IntCounter{}

	generator.On("Generate", "foo").Return("bar").Maybe()
	storage.On("LoadOrStore", "foo", "bar").Return("bar", nil).Maybe()

	zorro := New(generator, storage).WithMetrics(&Metrics{MaskOps: counter})
	counter.On("Add", int64(1)).Once()
	zorro.Mask("foo")

	counter.AssertExpectations(t)
}

func TestUnmaskIncrementsUnmaskOpCounter(t *testing.T) {
	storage := &storageMocks.Storage{}
	counter := &metricsMocks.IntCounter{}

	storage.On("Resolve", "foo").Return("bar", nil).Maybe()

	zorro := New(&generatorMocks.Generator{}, storage).WithMetrics(&Metrics{UnmaskOps: counter})
	counter.On("Add", int64(1)).Once()
	zorro.Unmask("foo")

	counter.AssertExpectations(t)
}
