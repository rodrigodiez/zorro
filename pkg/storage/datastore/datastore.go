package datastore

import (
	"cloud.google.com/go/datastore"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"golang.org/x/net/context"
)

// Item represents a Google Cloud Datastore entity
type Item struct {
	Data string
}

// Transaction represents a Google Cloud Datastore transaction
type Transaction interface {
	Get(*datastore.Key, interface{}) error
	Put(*datastore.Key, interface{}) (*datastore.PendingKey, error)
}

// TranslatorClient is a wrapper necessary to write tests around transactions
type TranslatorClient interface {
	RunInTransaction(context.Context, func(Transaction) error, ...datastore.TransactionOption) (*datastore.Commit, error)
	Get(context.Context, *datastore.Key, interface{}) error
}

type translator struct {
	client *datastore.Client
}

// NewTranslator returns a TranslatorClient
func NewTranslator(client *datastore.Client) TranslatorClient {
	return translator{client: client}
}

func (t translator) Get(ctx context.Context, key *datastore.Key, dest interface{}) error {
	return t.client.Get(ctx, key, dest)
}

func (t translator) RunInTransaction(ctx context.Context, f func(Transaction) error, opts ...datastore.TransactionOption) (*datastore.Commit, error) {
	return t.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		return f(tx)
	}, opts...)
}

type datastoreStorage struct {
	client    TranslatorClient
	keyKind   string
	valueKind string
}

func (d *datastoreStorage) LoadOrStore(key string, value string) (string, bool) {

	var (
		item        Item
		actualValue string
		loaded      bool
	)

	gKey := datastore.NameKey(d.keyKind, key, nil)

	d.client.RunInTransaction(context.TODO(), func(tx Transaction) error {

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

	if err := d.client.Get(context.TODO(), gKey, &item); err != nil {
		return "", false
	}

	return item.Data, true
}

func (d *datastoreStorage) Close() {
}

func (d *datastoreStorage) WithMetrics(metrics *storage.Metrics) storage.Storage {
	return d
}

// New creates a Storage persisted in Google Cloud Datastore
func New(client TranslatorClient, keyKind string, valueKind string) storage.Storage {
	return &datastoreStorage{client: client, keyKind: keyKind, valueKind: valueKind}
}
