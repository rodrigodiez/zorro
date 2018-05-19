package zorro

import (
	"github.com/boltdb/bolt"
)

type boltdb struct {
	db          *bolt.DB
	kBucketName []byte
	vBucketName []byte
}

// NewBoltDBStorage creates and initialises a new StorageCloser persisted in Bolt.
func NewBoltDBStorage(path string) (StorageCloser, error) {
	db, err := bolt.Open(path, 0600, nil)

	if err != nil {
		return nil, err
	}

	b := &boltdb{db: db, kBucketName: []byte("keys"), vBucketName: []byte("values")}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket(b.kBucketName)
		tx.CreateBucket(b.vBucketName)

		return nil
	})

	return b, nil
}

func (b *boltdb) Close() {
	b.db.Close()
}

func (b *boltdb) LoadOrStore(key string, value string) (actualMask string, loaded bool) {

	b.db.Update(func(tx *bolt.Tx) error {
		iBucket := tx.Bucket(b.kBucketName)
		mBucket := tx.Bucket(b.vBucketName)

		valueBytes := iBucket.Get([]byte(key))

		if valueBytes == nil {
			iBucket.Put([]byte(key), []byte(value))
			mBucket.Put([]byte(value), []byte(key))

			actualMask = value
			loaded = false

			return nil
		}

		valueBytesCopy := make([]byte, len(valueBytes))
		copy(valueBytesCopy, valueBytes)
		actualMask = string(valueBytesCopy)
		loaded = true

		return nil
	})

	return actualMask, loaded
}

func (b *boltdb) Resolve(value string) (key string, ok bool) {
	b.db.View(func(tx *bolt.Tx) error {
		mBucket := tx.Bucket(b.vBucketName)

		keyBytes := mBucket.Get([]byte(value))

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
