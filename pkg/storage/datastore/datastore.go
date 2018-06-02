package datastore

import (
	"cloud.google.com/go/datastore"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"golang.org/x/net/context"
)

type Client interface {
	RunInTransaction(context.Context, func(*datastore.Transaction) error, ...datastore.TransactionOption) (*datastore.Commit, error)
	Get(context.Context, *datastore.Key, interface{}) error
}

type Transaction interface {
	Get(*datastore.Key, interface{}) error
	Put(*datastore.Key, interface{}) (*datastore.PendingKey, error)
}

type datastoreStorage struct {
	cli       Client
	keyKind   string
	valueKind string
}

type Item struct {
	Data string
}

func (d *datastoreStorage) LoadOrStore(key string, value string) (string, bool) {

	var (
		item        Item
		actualValue string
		loaded      bool
	)

	gKey := datastore.NameKey(d.keyKind, key, nil)

	d.cli.RunInTransaction(context.TODO(), func(tx *datastore.Transaction) error {

		if err := tx.Get(gKey, &item); err != datastore.ErrNoSuchEntity {
			actualValue = item.Data
			loaded = true

			return err
		}

		tx.Put(gKey, &Item{Data: value})
		tx.Put(datastore.NameKey(d.valueKind, value, nil), &Item{Data: key})

		actualValue = value
		loaded = false

		return nil
	})

	return actualValue, loaded
}

func (d *datastoreStorage) Resolve(value string) (string, bool) {
	var item Item
	gKey := datastore.NameKey(d.valueKind, value, nil)

	if err := d.cli.Get(context.TODO(), gKey, &item); err != nil {
		return "", false
	}

	return item.Data, true
}

func (d *datastoreStorage) Close() {
}

func (d *datastoreStorage) WithMetrics(metrics *storage.Metrics) storage.Storage {
	return d
}

func New(client Client, keyKind string, valueKind string) storage.Storage {
	return &datastoreStorage{cli: client, keyKind: keyKind, valueKind: valueKind}
}
