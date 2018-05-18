package zorro

import (
	"github.com/satori/go.uuid"
)

type uuidv4Generator struct {
}

func (*uuidv4Generator) Generate(_ string) (mask string) {
	return uuid.NewV4().String()
}

// NewUUIDv4Generator creates a new Generator that generates UUIDv4 masks.
//
// See: https://en.wikipedia.org/wiki/Universally_unique_identifier
func NewUUIDv4Generator() Generator {
	return &uuidv4Generator{}
}
