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
	svc *datastore.Client
}

// NewTranslator returns a TranslatorClient
func NewTranslator(svc *datastore.Client) TranslatorClient {
	return translator{svc: svc}
}

func (t translator) Get(ctx context.Context, key *datastore.Key, dest interface{}) error {
	return t.svc.Get(ctx, key, dest)
}

func (t translator) RunInTransaction(ctx context.Context, f func(Transaction) error, opts ...datastore.TransactionOption) (*datastore.Commit, error) {
	return t.svc.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		return f(tx)
	}, opts...)
}

type datastoreStorage struct {
	client    TranslatorClient
	keyKind   string
	valueKind string
}

// New creates a Storage persisted in Google Cloud Datastore
func New(client TranslatorClient, keyKind string, valueKind string) storage.Storage {
	return &datastoreStorage{client: client, keyKind: keyKind, valueKind: valueKind}
}

func (sto *datastoreStorage) LoadOrStore(key string, value string) (string, error) {

	var (
		item        Item
		err         error
		actualValue string
	)

	gKey := datastore.NameKey(sto.keyKind, key, nil)

	_, err = sto.client.RunInTransaction(context.TODO(), func(tx Transaction) error {
		var err error

		if err = tx.Get(gKey, &item); err != datastore.ErrNoSuchEntity {
			actualValue = item.Data

			return nil
		}

		if err != nil {
			return err
		}

		_, err = tx.Put(gKey, &Item{Data: value})
		if err != nil {
			return err
		}
		tx.Put(datastore.NameKey(sto.valueKind, value, nil), &Item{Data: key})

		actualValue = value

		return nil
	})

	return actualValue, err
}

func (sto *datastoreStorage) Resolve(value string) (string, error) {
	var item Item
	gKey := datastore.NameKey(sto.valueKind, value, nil)

	if err := sto.client.Get(context.TODO(), gKey, &item); err != nil {
		return "", err
	}

	return item.Data, nil
}

// Close is noop
func (sto *datastoreStorage) Close() {
}

func (sto *datastoreStorage) WithMetrics(metrics *storage.Metrics) storage.Storage {
	return sto
}
