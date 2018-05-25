package boltdb

import (
	"github.com/boltdb/bolt"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

type boltdb struct {
	db           *bolt.DB
	keysBucket   []byte
	valuesBucket []byte
	metrics      *storage.Metrics
}

// New creates and initialises a new Closer persisted in Bolt.
func New(path string) (storage.Storage, error) {
	db, err := bolt.Open(path, 0600, nil)

	if err != nil {
		return nil, err
	}

	b := &boltdb{db: db, keysBucket: []byte("keys"), valuesBucket: []byte("values")}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket(b.keysBucket)
		tx.CreateBucket(b.valuesBucket)

		return nil
	})

	return b, nil
}

func (b *boltdb) Close() {
	b.db.Close()
}

func (b *boltdb) LoadOrStore(key string, value string) (actualMask string, loaded bool) {

	b.db.Update(func(tx *bolt.Tx) error {
		keysBucket := tx.Bucket(b.keysBucket)
		valuesBucket := tx.Bucket(b.valuesBucket)

		valueBytes := keysBucket.Get([]byte(key))

		if valueBytes == nil {
			keysBucket.Put([]byte(key), []byte(value))
			valuesBucket.Put([]byte(value), []byte(key))

			b.incrStoreOps()

			actualMask = value
			loaded = false

			return nil
		}

		valueBytesCopy := make([]byte, len(valueBytes))
		copy(valueBytesCopy, valueBytes)

		b.incrLoadOps()

		actualMask = string(valueBytesCopy)
		loaded = true

		return nil
	})

	return actualMask, loaded
}

func (b *boltdb) Resolve(value string) (key string, ok bool) {
	b.db.View(func(tx *bolt.Tx) error {
		valuesBucket := tx.Bucket(b.valuesBucket)

		keyBytes := valuesBucket.Get([]byte(value))
		b.incrResolveOps()

		if keyBytes == nil {
			ok = false
			key = ""

			return nil
		}

		keyBytesCopy := make([]byte, len(keyBytes))
		copy(keyBytesCopy, keyBytes)
		key = string(keyBytesCopy)
		ok = true

		return nil
	})

	return key, ok
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
