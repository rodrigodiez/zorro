package zorro

import (
	"github.com/boltdb/bolt"
)

type boltdb struct {
	db          *bolt.DB
	iBucketName []byte
	mBucketName []byte
}

// NewBoltDBStorage creates and initialises a new StorageCloser persisted in Bolt.
func NewBoltDBStorage(path string) (StorageCloser, error) {
	db, err := bolt.Open(path, 0600, nil)

	if err != nil {
		return nil, err
	}

	b := &boltdb{db: db, iBucketName: []byte("ids"), mBucketName: []byte("masks")}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket(b.iBucketName)
		tx.CreateBucket(b.mBucketName)

		return nil
	})

	return b, nil
}

func (b *boltdb) Close() {
	b.db.Close()
}

func (b *boltdb) LoadOrStore(id string, mask string) (actualMask string, loaded bool) {

	b.db.Update(func(tx *bolt.Tx) error {
		iBucket := tx.Bucket(b.iBucketName)
		mBucket := tx.Bucket(b.mBucketName)

		maskBytes := iBucket.Get([]byte(id))

		if maskBytes == nil {
			iBucket.Put([]byte(id), []byte(mask))
			mBucket.Put([]byte(mask), []byte(id))

			actualMask = mask
			loaded = false

			return nil
		}

		maskBytesCopy := make([]byte, len(maskBytes))
		copy(maskBytesCopy, maskBytes)
		actualMask = string(maskBytesCopy)
		loaded = true

		return nil
	})

	return actualMask, loaded
}

func (b *boltdb) Resolve(mask string) (id string, ok bool) {
	b.db.View(func(tx *bolt.Tx) error {
		mBucket := tx.Bucket(b.mBucketName)

		idBytes := mBucket.Get([]byte(mask))

		if idBytes == nil {
			ok = false
			id = ""

			return nil
		}

		idBytesCopy := make([]byte, len(idBytes))
		copy(idBytesCopy, idBytes)
		id = string(idBytesCopy)
		ok = true

		return nil
	})

	return id, ok
}
