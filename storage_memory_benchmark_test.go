package zorro

import (
	"math/rand"
	"testing"
)

func BenchmarkMemory(b *testing.B) {
	storage := NewInMemoryStorage()

	for i := 0; i < b.N; i++ {
		key := randStringBytes(24)
		value := randStringBytes(24)

		storage.LoadOrStore(key, value)
	}
}

func BenchmarkMemoryParallel(b *testing.B) {
	storage := NewInMemoryStorage()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := randStringBytes(24)
			value := randStringBytes(24)

			storage.LoadOrStore(key, value)
		}
	})
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
