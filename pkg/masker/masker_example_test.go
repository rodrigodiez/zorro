package masker

import (
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

func ExampleNew() {
	masker := New(generator.NewUUIDv4(), storage.NewMem())
	masker.Mask("foo")
}
