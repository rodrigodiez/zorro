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

// ClientAdapter is a wrapper necessary to write tests based on interfaces
type ClientAdapter interface {
	RunInTransaction(context.Context, func(Transaction) error, ...datastore.TransactionOption) (*datastore.Commit, error)
	Get(context.Context, *datastore.Key, interface{}) error
}

type adapter struct {
	svc *datastore.Client
}

// NewClientAdapter returns a ClientAdapter
func NewClientAdapter(svc *datastore.Client) ClientAdapter {
	return adapter{svc: svc}
}

func (t adapter) Get(ctx context.Context, key *datastore.Key, dest interface{}) error {
	return t.svc.Get(ctx, key, dest)
}

func (t adapter) RunInTransaction(ctx context.Context, f func(Transaction) error, opts ...datastore.TransactionOption) (*datastore.Commit, error) {
	return t.svc.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		return f(tx)
	}, opts...)
}

type datastoreStorage struct {
	adapter   ClientAdapter
	keyKind   string
	valueKind string
}

// New creates a Storage persisted in Google Cloud Datastore
func New(adapter ClientAdapter, keyKind string, valueKind string) storage.Storage {
	return &datastoreStorage{adapter: adapter, keyKind: keyKind, valueKind: valueKind}
}

func (sto *datastoreStorage) LoadOrStore(key string, value string) (string, error) {

	var (
		item        Item
		err         error
		actualValue string
	)

	_, err = sto.adapter.RunInTransaction(context.TODO(), func(tx Transaction) error {

		gKey := datastore.NameKey(sto.keyKind, key, nil)
		err = tx.Get(gKey, &item)

		switch {
		case err == nil:
			actualValue = item.Data

			return nil
		case err != datastore.ErrNoSuchEntity:
			return err
		default:
		}

		_, err = tx.Put(gKey, &Item{Data: value})
		if err != nil {
			return err
		}

		_, err := tx.Put(datastore.NameKey(sto.valueKind, value, nil), &Item{Data: key})
		if err != nil {
			return err
		}

		actualValue = value

		return nil
	})

	if err != nil {
		return "", err
	}

	return actualValue, err
}

func (sto *datastoreStorage) Resolve(value string) (string, error) {
	var item Item
	gKey := datastore.NameKey(sto.valueKind, value, nil)

	if err := sto.adapter.Get(context.TODO(), gKey, &item); err != nil {
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
