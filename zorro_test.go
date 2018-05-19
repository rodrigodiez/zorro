package zorro

import (
	"testing"

	"github.com/rodrigodiez/zorro/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewReturnsZorro(t *testing.T) {
	t.Parallel()

	var _ Zorro = New(&mocks.Generator{}, &mocks.Storage{})
}

func TestMask(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		key    string
		value  string
		actual string
		loaded bool
	}{
		{name: "value was not loaded by storage", key: "foo", value: "bar", actual: "bar", loaded: false},
		{name: "value was loaded by storage", key: "foo", value: "bar", actual: "baz", loaded: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			generator := &mocks.Generator{}
			storage := &mocks.Storage{}

			zorro := New(generator, storage)

			generator.On("Generate", tc.key).Return(tc.value).Once()
			storage.On("LoadOrStore", tc.key, tc.value).Return(tc.actual, tc.loaded).Once()

			value := zorro.Mask(tc.key)
			assert.Equal(t, tc.actual, value)

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
		ok    bool
	}{
		{name: "value was unmasked", value: "bar", key: "foo", ok: true},
		{name: "value was not found", value: "bar", key: "", ok: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			storage := &mocks.Storage{}

			zorro := New(&mocks.Generator{}, storage)

			storage.On("Resolve", tc.value).Return(tc.key, tc.ok).Once()

			key, ok := zorro.Unmask(tc.value)

			assert.Equal(t, tc.key, key)
			assert.Equal(t, tc.ok, ok)

			storage.AssertExpectations(t)
		})
	}
}
