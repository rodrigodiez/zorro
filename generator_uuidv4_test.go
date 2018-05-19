package zorro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsGenerator(t *testing.T) {
	t.Parallel()
	var _ Generator = NewUUIDv4Generator()
}

func TestGenerateReturnsRandomString(t *testing.T) {
	t.Parallel()

	gen := NewUUIDv4Generator()

	firstValue := gen.Generate("")
	secondValue := gen.Generate("")

	assert.NotEqual(t, firstValue, secondValue)
}
