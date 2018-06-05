package datastore_test

import (
	"errors"
	"testing"

	goDatastore "cloud.google.com/go/datastore"
	datastoreMocks "github.com/rodrigodiez/zorro/lib/mocks/storage/datastore"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/rodrigodiez/zorro/pkg/storage/datastore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImplementsStorage(t *testing.T) {
	var _ storage.Storage = datastore.New(&datastoreMocks.TranslatorClient{}, "keyKind", "value")
}

func TestLoadOrStoreReturnsValueAndNilIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	translator := &datastoreMocks.TranslatorClient{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.Anything, mock.Anything).Return(goDatastore.ErrNoSuchEntity).Maybe()

	tx.On("Put", mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.MatchedBy(func(item *datastore.Item) bool {
		return item.Data == "bar"
	})).Return(&goDatastore.PendingKey{}, nil).Once()

	tx.On("Put", mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.MatchedBy(func(item *datastore.Item) bool {
		return item.Data == "foo"
	})).Return(&goDatastore.PendingKey{}, nil).Once()

	sto := datastore.New(translator, "keyKind", "valueKind")

	translator.On("RunInTransaction", mock.Anything, mock.Anything).Return(&goDatastore.Commit{}, nil).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(datastore.Transaction) error)

		f(tx)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
	tx.AssertExpectations(t)
	translator.AssertExpectations(t)
}

func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
	t.Parallel()

	translator := &datastoreMocks.TranslatorClient{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		item := args.Get(1).(*datastore.Item)

		item.Data = "baz"
	}).Once()

	sto := datastore.New(translator, "keyKind", "valueKind")

	translator.On("RunInTransaction", mock.Anything, mock.Anything).Return(&goDatastore.Commit{}, nil).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx datastore.Transaction) error)

		f(tx)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "baz", value)
	assert.Nil(t, err)
	tx.AssertExpectations(t)
	translator.AssertExpectations(t)
}

func TestLoadOrStoreReturnsEmptyStringAndErrorIfStorageFailsGetting(t *testing.T) {
	t.Parallel()

	translator := &datastoreMocks.TranslatorClient{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil).Once()

	sto := datastore.New(translator, "keyKind", "valueKind")

	translator.On("RunInTransaction", mock.Anything, mock.Anything).Return(&goDatastore.Commit{}, errors.New("")).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx datastore.Transaction) error)

		f(tx)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "", value)
	assert.NotNil(t, err)
	tx.AssertExpectations(t)
	translator.AssertExpectations(t)
}

func TestLoadOrStoreReturnsEmptyStringAndErrorIfStorageFailsPutting(t *testing.T) {
	t.Parallel()

	translator := &datastoreMocks.TranslatorClient{}
	tx := &datastoreMocks.Transaction{}

	tx.On("Get", mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		item := args.Get(1).(*datastore.Item)

		item.Data = "baz"
	}).Once()

	tx.On("Put", mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "keyKind" && gKey.Name == "foo"
	}), mock.Anything).Return(nil, errors.New("")).Once()

	sto := datastore.New(translator, "keyKind", "valueKind")

	translator.On("RunInTransaction", mock.Anything, mock.Anything).Return(&goDatastore.Commit{}, errors.New("")).Run(func(args mock.Arguments) {
		f := args.Get(1).(func(tx datastore.Transaction) error)

		f(tx)
	}).Once()

	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "", value)
	assert.NotNil(t, err)
	tx.AssertExpectations(t)
	translator.AssertExpectations(t)
}

func TestResolveReturnsKeyAndNilIfValueFound(t *testing.T) {
	t.Parallel()

	translator := &datastoreMocks.TranslatorClient{}

	translator.On("Get", mock.Anything, mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		item := args.Get(2).(*datastore.Item)

		item.Data = "foo"
	}).Once()

	sto := datastore.New(translator, "keyKind", "valueKind")

	key, err := sto.Resolve("bar")

	assert.Equal(t, "foo", key)
	assert.Nil(t, err)
	translator.AssertExpectations(t)
}

func TestResolveReturnsEmptyStringAndErrorIfValueNotFound(t *testing.T) {
	t.Parallel()

	translator := &datastoreMocks.TranslatorClient{}

	translator.On("Get", mock.Anything, mock.MatchedBy(func(gKey *goDatastore.Key) bool {
		return gKey.Kind == "valueKind" && gKey.Name == "bar"
	}), mock.Anything).Return(goDatastore.ErrNoSuchEntity).Once()

	sto := datastore.New(translator, "keyKind", "valueKind")

	key, err := sto.Resolve("bar")

	assert.Empty(t, key)
	assert.NotNil(t, err)
	translator.AssertExpectations(t)
}
