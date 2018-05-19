package zorro

import (
	"os"
	"testing"
)

func BenchmarkBoltDb(b *testing.B) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	defer storage.Close()

	for i := 0; i < b.N; i++ {
		key := randStringBytes(24)
		value := randStringBytes(24)

		storage.LoadOrStore(key, value)
	}
}

func BenchmarkBoltDBParallel(b *testing.B) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	defer storage.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := randStringBytes(24)
			value := randStringBytes(24)

			storage.LoadOrStore(key, value)
		}
	})
}
