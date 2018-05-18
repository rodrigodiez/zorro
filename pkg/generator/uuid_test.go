package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsGenerator(t *testing.T) {
	var _ Generator = NewUUIDv4()
}

func TestGenerateReturnsRandomString(t *testing.T) {
	gen := NewUUIDv4()

	firstValue := gen.Generate("")
	secondValue := gen.Generate("")

	assert.NotEqual(t, firstValue, secondValue)
}
