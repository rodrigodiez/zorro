package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodrigodiez/zorro/lib/mocks"
)

func TestNewCallMetricsImplementsMiddleWare(t *testing.T) {
	var _ MiddleWare = NewCallMetrics(&mocks.Counter{})
}

func TestCallMetricsIncreasesMaskCountOnMask(t *testing.T) {
	counter := &mocks.Counter{}
	zorro := &mocks.Zorro{}
	middleware := NewCallMetrics(counter)
	wrapped := middleware(zorro)

	zorro.On("Mask", "foo").Return("bar").Maybe()
	counter.On("Add", int64(1)).Once()
	wrapped.Mask("foo")

	counter.AssertExpectations(t)
}
func TestCallMetricsForwardsToNext(t *testing.T) {
	counter := &mocks.Counter{}
	zorro := &mocks.Zorro{}
	middleware := NewCallMetrics(counter)
	wrapped := middleware(zorro)

	zorro.On("Mask", "foo").Return("bar").Once()
	counter.On("Add", int64(1)).Maybe()

	value := wrapped.Mask("foo")

	assert.Equal(t, "bar", value)
	zorro.AssertExpectations(t)
}
