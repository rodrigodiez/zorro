package storage

import (
	"math/rand"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	mem := NewMem()

	for i := 0; i < b.N; i++ {
		id := RandStringBytes(24)
		mask := RandStringBytes(24)

		mem.LoadOrStore(id, mask)
	}
}

func BenchmarkSetParallel(b *testing.B) {
	mem := NewMem()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			id := RandStringBytes(24)
			mask := RandStringBytes(24)

			mem.LoadOrStore(id, mask)
		}
	})
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
