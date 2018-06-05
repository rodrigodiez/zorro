package memory

import (
	"testing"

	"github.com/rodrigodiez/zorro/lib/random"
)

func BenchmarkMemory(b *testing.B) {
	sto := New()

	for i := 0; i < b.N; i++ {
		key := random.NewString(24)
		value := random.NewString(24)

		sto.LoadOrStore(key, value)
	}
}

func BenchmarkMemoryParallel(b *testing.B) {
	sto := New()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := random.NewString(24)
			value := random.NewString(24)

			sto.LoadOrStore(key, value)
		}
	})
}
