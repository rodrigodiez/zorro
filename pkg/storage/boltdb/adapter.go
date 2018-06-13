package boltdb

import "github.com/boltdb/bolt"

// Transaction represents a BoltDB transaction
type Transaction interface {
	CreateBucket([]byte) (*bolt.Bucket, error)
	Bucket([]byte) *bolt.Bucket
}

// Bucket represents a BoltDB bucket
type Bucket interface {
	Put([]byte, []byte) error
	Get(key []byte) []byte
}

// ClientAdapter is a wrapper necessary to write tests based on interfaces
type ClientAdapter interface {
	Update(func(tx Transaction) error) error
	View(func(tx Transaction) error) error
	Close() error
}

type adapter struct {
	svc *bolt.DB
}

// NewClientAdapter returns a ClientAdapter
func NewClientAdapter(db *bolt.DB) ClientAdapter {

	return &adapter{svc: db}
}

func (adapter *adapter) Close() error {
	return adapter.svc.Close()
}

func (adapter *adapter) Update(f func(tx Transaction) error) error {
	return adapter.svc.Update(func(tx *bolt.Tx) error {
		return f(tx)
	})
}

func (adapter *adapter) View(f func(tx Transaction) error) error {
	return adapter.svc.View(func(tx *bolt.Tx) error {
		return f(tx)
	})
}
