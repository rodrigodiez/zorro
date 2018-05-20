package service

import (
	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func ExampleNew() {
	zorro := New(uuid.NewV4(), memory.New())
	zorro.Mask("foo")
}
