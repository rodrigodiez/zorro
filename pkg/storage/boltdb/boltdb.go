package boltdb

import (
	"github.com/rodrigodiez/zorro/pkg/storage"
)

type boltdb struct {
	adapter      ClientAdapter
	keysBucket   []byte
	valuesBucket []byte
	metrics      *storage.Metrics
}

// New creates and initialises a new Closer persisted in Bolt.
func New(adapter ClientAdapter) (storage.Storage, error) {

	b := &boltdb{adapter: adapter, keysBucket: []byte("keys"), valuesBucket: []byte("values")}

	err := adapter.Update(func(tx Transaction) error {
		tx.CreateBucket(b.keysBucket)
		tx.CreateBucket(b.valuesBucket)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return b, nil

}

// Close closes the underlying storage
func (b *boltdb) Close() {
	b.adapter.Close()
}

func (b *boltdb) LoadOrStore(key string, value string) (string, error) {

	var actual string

	b.adapter.Update(func(tx Transaction) error {
		keysBucket := tx.Bucket(b.keysBucket)
		valuesBucket := tx.Bucket(b.valuesBucket)

		valueBytes := keysBucket.Get([]byte(key))

		if valueBytes == nil {
			keysBucket.Put([]byte(key), []byte(value))
			valuesBucket.Put([]byte(value), []byte(key))

			b.incrStoreOps()

			actual = value

			return nil
		}

		valueBytesCopy := make([]byte, len(valueBytes))
		copy(valueBytesCopy, valueBytes)

		b.incrLoadOps()

		actual = string(valueBytesCopy)

		return nil
	})

	return actual, nil
}

func (b *boltdb) Resolve(value string) (string, error) {
	var key string

	b.adapter.View(func(tx Transaction) error {
		valuesBucket := tx.Bucket(b.valuesBucket)

		keyBytes := valuesBucket.Get([]byte(value))
		b.incrResolveOps()

		if keyBytes == nil {
			key = ""

			return nil
		}

		keyBytesCopy := make([]byte, len(keyBytes))
		copy(keyBytesCopy, keyBytes)
		key = string(keyBytesCopy)

		return nil
	})

	return key, nil
}

func (b *boltdb) WithMetrics(metrics *storage.Metrics) storage.Storage {
	b.metrics = metrics

	return b
}

func (b *boltdb) incrStoreOps() {
	if b.metrics != nil && b.metrics.StoreOps != nil {
		b.metrics.StoreOps.Add(int64(1))
	}
}

func (b *boltdb) incrLoadOps() {
	if b.metrics != nil && b.metrics.LoadOps != nil {
		b.metrics.LoadOps.Add(int64(1))
	}
}

func (b *boltdb) incrResolveOps() {
	if b.metrics != nil && b.metrics.ResolveOps != nil {
		b.metrics.ResolveOps.Add(int64(1))
	}
}
