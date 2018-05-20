package uuid

import (
	"testing"

	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func TestImplementsGenerator(t *testing.T) {
	t.Parallel()
	var _ generator.Generator = NewV4()
}

func TestGenerateReturnsRandomString(t *testing.T) {
	t.Parallel()

	gen := NewV4()

	firstValue := gen.Generate("")
	secondValue := gen.Generate("")

	assert.NotEqual(t, firstValue, secondValue)
}
