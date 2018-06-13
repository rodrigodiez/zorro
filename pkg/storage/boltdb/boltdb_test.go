package boltdb_test

import (
	"testing"

	boltMocks "github.com/rodrigodiez/zorro/lib/mocks/storage/boltdb"
	"github.com/rodrigodiez/zorro/pkg/storage"
	zorroBolt "github.com/rodrigodiez/zorro/pkg/storage/boltdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImplementsStorage(t *testing.T) {
	t.Parallel()

	adapter := &boltMocks.ClientAdapter{}
	adapter.On("Update", mock.Anything).Return(nil).Maybe()

	sto, _ := zorroBolt.New(adapter)

	func(sto storage.Storage) {
		return
	}(sto)
}

func TestCloseClosesTheDB(t *testing.T) {
	t.Parallel()

	adapter := &boltMocks.ClientAdapter{}
	adapter.On("Update", mock.Anything).Return(nil).Maybe()
	adapter.On("Close").Return(nil).Once()

	sto, _ := zorroBolt.New(adapter)

	sto.Close()
	adapter.AssertExpectations(t)
}

func TestNewCreatesKeysAndValuesBuckets(t *testing.T) {
	t.Parallel()

	adapter := &boltMocks.ClientAdapter{}
	tx := &boltMocks.Transaction{}

	tx.On("CreateBucket", mock.Anything).Return(nil, nil).Maybe()

	adapter.On("Update", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(zorroBolt.Transaction) error)

		err := f(tx)
		assert.Nil(t, err)
	})

	zorroBolt.New(adapter)

	tx.AssertCalled(t, "CreateBucket", []byte("keys"))
	tx.AssertCalled(t, "CreateBucket", []byte("values"))
}

func TestLoadOrStoreReTestturnsValueAndNilIfIdDoesNotExist(t *testing.T) {
	t.Parallel()

	adapter := &boltMocks.ClientAdapter{}
	tx := &boltMocks.Transaction{}
	keysBucket := &boltMocks.Bucket{}
	valuesBucket := &boltMocks.Bucket{}

	adapter.On("Update", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(zorroBolt.Transaction) error)

		err := f(tx)
		assert.Nil(t, err)
	})

	tx.On("CreateBucket", mock.Anything).Return(nil, nil).Maybe()
	tx.On("Bucket", []byte("keys")).Return(keysBucket).Once()
	tx.On("Bucket", []byte("values")).Return(valuesBucket).Once()

	keysBucket.On("Get", []byte("foo")).Return(nil).Once()
	keysBucket.On("Put", mock.Anything).Return(nil).Maybe()
	valuesBucket.On("Put", mock.Anything).Return(nil).Maybe()

	sto, _ := zorroBolt.New(adapter)
	value, err := sto.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

// func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
// 	t.Parallel()

// 	path := getTmpPath()
// 	defer os.Remove(path)

// 	storage, _ := zorroBolt.New(path)
// 	defer storage.Close()

// 	storage.LoadOrStore("foo", "bar")
// 	value, err := storage.LoadOrStore("foo", "baz")

// 	assert.Equal(t, "bar", value)
// 	assert.Nil(t, err)
// }

// func TestLoadOrStoreReturnsEmptyStringAndErrIfStorageFails(t *testing.T) {
// 	t.Parallel()

// 	path := getTmpPath()
// 	defer os.Remove(path)

// 	storage, _ := zorroBolt.New(path)
// 	defer storage.Close()

// 	storage.LoadOrStore("foo", "bar")
// 	value, err := storage.LoadOrStore("foo", "baz")

// 	assert.Equal(t, "bar", value)
// 	assert.Nil(t, err)
// }

// func TestResolve(t *testing.T) {
// 	t.Parallel()

// 	tt := []struct {
// 		name        string
// 		loadedID    string
// 		loadedValue string
// 		value       string
// 		expectedID  string
// 		expectedOk  bool
// 	}{
// 		{name: "Id exists", loadedID: "foo", loadedValue: "bar", value: "bar", expectedID: "foo", expectedOk: true},
// 		{name: "Id does not exist", loadedID: "foo", loadedValue: "bar", value: "baz", expectedID: "", expectedOk: false},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			path := getTmpPath()
// 			defer os.Remove(path)

// 			storage, _ := zorroBolt.New(path)
// 			defer storage.Close()

// 			storage.LoadOrStore(tc.loadedID, tc.loadedValue)
// 			key, ok := storage.Resolve(tc.value)

// 			assert.Equal(t, tc.expectedID, key)
// 			assert.Equal(t, tc.expectedOk, ok)
// 		})
// 	}
// }

// func TestLoadOrStoreIncrementsStoreOpsIfKeyDoesNotExist(t *testing.T) {
// 	t.Parallel()

// 	counter := &metricsMocks.IntCounter{}
// 	counter.On("Add", int64(1))

// 	path := getTmpPath()
// 	defer os.Remove(path)

// 	bolt, _ := zorroBolt.New(path)
// 	bolt.WithMetrics(&storage.Metrics{StoreOps: counter})
// 	defer bolt.Close()

// 	bolt.LoadOrStore("foo", "bar")

// 	counter.AssertCalled(t, "Add", int64(1))
// }

// func TestLoadOrStoreIncrementsLoadOpsIfKeyExists(t *testing.T) {
// 	t.Parallel()

// 	counter := &metricsMocks.IntCounter{}
// 	counter.On("Add", int64(1))

// 	path := getTmpPath()
// 	defer os.Remove(path)

// 	bolt, _ := zorroBolt.New(path)
// 	bolt.WithMetrics(&storage.Metrics{LoadOps: counter})
// 	defer bolt.Close()

// 	bolt.LoadOrStore("foo", "bar")
// 	bolt.LoadOrStore("foo", "bar")

// 	counter.AssertCalled(t, "Add", int64(1))
// }
// func TestResolveIncrementsResolveOps(t *testing.T) {
// 	t.Parallel()

// 	counter := &metricsMocks.IntCounter{}
// 	counter.On("Add", int64(1))

// 	path := getTmpPath()
// 	defer os.Remove(path)

// 	bolt, _ := zorroBolt.New(path)
// 	bolt.WithMetrics(&storage.Metrics{ResolveOps: counter})
// 	defer bolt.Close()

// 	bolt.Resolve("bar")

// 	counter.AssertCalled(t, "Add", int64(1))
// }

// func getTmpPath() string {
// 	f, err := ioutil.TempFile("", "zorro-tests")

// 	if err != nil {
// 		panic(err)
// 	}

// 	path := f.Name()
// 	defer os.Remove(path)

// 	f.Close()

// 	return path
// }
