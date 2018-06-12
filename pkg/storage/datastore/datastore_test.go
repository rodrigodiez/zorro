package datastore_test

import (
	"errors"
	"testing"

	"cloud.google.com/go/datastore"
	datastoreMocks "github.com/rodrigodiez/zorro/lib/mocks/storage/datastore"
	"github.com/rodrigodiez/zorro/pkg/storage"
	zorroDatastore "github.com/rodrigodiez/zorro/pkg/storage/datastore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImplementsStorage(t *testing.T) {
	var _ storage.Storage = zorroDatastore.New(&datastoreMocks.ClientAdapter{}, "keyKind", "value")
}

func TestLoadOrStoreReturnsValueAndNilIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.Anything, mock.Anything).Return(datastore.ErrNoSuchEntity).Maybe()

	tx.On("Put", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.MatchedBy(func(item *zorroDatastore.Item) bool {
		return item.Data == "bar"
	})).Return(&datastore.PendingKey{}, nil).Once()

	tx.On("Put", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.MatchedBy(func(item *zorroDatastore.Item) bool {
		return item.Data == "foo"
	})).Return(&datastore.PendingKey{}, nil).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	adapter.On("RunInTransaction", mock.Anything, mock.Anything).Return(&datastore.Commit{}, nil).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(zorroDatastore.Transaction) error)

		err := f(tx)

		assert.Nil(t, err)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
	tx.AssertExpectations(t)
	adapter.AssertExpectations(t)
}

func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		item := args.Get(1).(*zorroDatastore.Item)

		item.Data = "baz"
	}).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	adapter.On("RunInTransaction", mock.Anything, mock.Anything).Return(&datastore.Commit{}, nil).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx zorroDatastore.Transaction) error)

		err := f(tx)

		assert.Nil(t, err)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "baz", value)
	assert.Nil(t, err)
	tx.AssertExpectations(t)
	adapter.AssertExpectations(t)
}

func TestLoadOrStoreReturnsEmptyStringAndErrorIfStorageFailsGetting(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(errors.New("")).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	adapter.On("RunInTransaction", mock.Anything, mock.Anything).Return(&datastore.Commit{}, errors.New("")).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx zorroDatastore.Transaction) error)

		err := f(tx)

		assert.NotNil(t, err)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "", value)
	assert.NotNil(t, err)
	tx.AssertExpectations(t)
	adapter.AssertExpectations(t)
}

func TestLoadOrStoreReturnsEmptyStringAndErrorIfStorageFailsPuttingKey(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(datastore.ErrNoSuchEntity).Once()

	tx.On("Put", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil, errors.New("")).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	adapter.On("RunInTransaction", mock.Anything, mock.Anything).Return(&datastore.Commit{}, errors.New("")).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx zorroDatastore.Transaction) error)

		err := f(tx)

		assert.NotNil(t, err)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "", value)
	assert.NotNil(t, err)
	tx.AssertExpectations(t)
	adapter.AssertExpectations(t)
}

func TestLoadOrStoreReturnsEmptyStringAndErrorIfStorageFailsPuttingValue(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(datastore.ErrNoSuchEntity).Once()

	tx.On("Put", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil, nil).Once()

	tx.On("Put", mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.Anything).Return(nil, errors.New("")).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	adapter.On("RunInTransaction", mock.Anything, mock.Anything).Return(&datastore.Commit{}, errors.New("")).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx zorroDatastore.Transaction) error)
		err := f(tx)

		assert.NotNil(t, err)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "", value)
	assert.NotNil(t, err)
	tx.AssertExpectations(t)
	adapter.AssertExpectations(t)
}

func TestResolveReturnsKeyAndNilIfValueFound(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}

	adapter.On("Get", mock.Anything, mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		item := args.Get(2).(*zorroDatastore.Item)

		item.Data = "foo"
	}).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	key, err := sto.Resolve("bar")

	assert.Equal(t, "foo", key)
	assert.Nil(t, err)
	adapter.AssertExpectations(t)
}

func TestResolveReturnsEmptyStringAndErrorIfValueNotFound(t *testing.T) {
	t.Parallel()

	adapter := &datastoreMocks.ClientAdapter{}

	adapter.On("Get", mock.Anything, mock.MatchedBy(func(gKey *datastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.Anything).Return(datastore.ErrNoSuchEntity).Once()

	sto := zorroDatastore.New(adapter, "keyKind", "valueKind")

	key, err := sto.Resolve("bar")

	assert.Empty(t, key)
	assert.NotNil(t, err)
	adapter.AssertExpectations(t)
}
