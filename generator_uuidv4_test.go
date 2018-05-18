package zorro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsGenerator(t *testing.T) {
	var _ Generator = NewUUIDv4Generator()
}

func TestGenerateReturnsRandomString(t *testing.T) {
	gen := NewUUIDv4Generator()

	firstValue := gen.Generate("")
	secondValue := gen.Generate("")

	assert.NotEqual(t, firstValue, secondValue)
}
