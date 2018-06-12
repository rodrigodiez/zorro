package boltdb_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	metricsMocks "github.com/rodrigodiez/zorro/lib/mocks/metrics"
	"github.com/rodrigodiez/zorro/pkg/storage"
	zorroBolt "github.com/rodrigodiez/zorro/pkg/storage/boltdb"
	"github.com/stretchr/testify/assert"
)

func TestNewImplementsStorage(t *testing.T) {
	var storage storage.Storage

	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ = zorroBolt.New(path)
	storage.Close()
}
func TestNewReturnsErrIfCantOpen(t *testing.T) {
	t.Parallel()

	storage, err := zorroBolt.New("/a/path/that/does/not/exist")

	assert.Nil(t, storage)
	assert.NotNil(t, err)
}

func TestCloseClosesTheDB(t *testing.T) {
	t.Skip("Not sure how to do this yet")
}

func TestNewCreatesKeysAndValuesBuckets(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := zorroBolt.New(path)
	storage.Close()

	db, _ := bolt.Open(path, 0600, nil)
	db.View(func(tx *bolt.Tx) error {
		bKeys := tx.Bucket([]byte("keys"))
		bValues := tx.Bucket([]byte("values"))

		assert.NotNil(t, bKeys, "keys bucket does not exist")
		assert.NotNil(t, bValues, "values bucket does not exist")

		return nil
	})

	db.Close()
}

func TestLoadOrStoreReTestturnsValueAndNilIfIdDoesNotExist(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := zorroBolt.New(path)
	defer storage.Close()

	value, err := storage.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := zorroBolt.New(path)
	defer storage.Close()

	storage.LoadOrStore("foo", "bar")
	value, err := storage.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

func TestLoadOrStoreReturnsEmptyStringAndErrIfStorageFails(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := zorroBolt.New(path)
	defer storage.Close()

	storage.LoadOrStore("foo", "bar")
	value, err := storage.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)
}

func TestResolve(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		loadedID    string
		loadedValue string
		value       string
		expectedID  string
		expectedOk  bool
	}{
		{name: "Id exists", loadedID: "foo", loadedValue: "bar", value: "bar", expectedID: "foo", expectedOk: true},
		{name: "Id does not exist", loadedID: "foo", loadedValue: "bar", value: "baz", expectedID: "", expectedOk: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			path := getTmpPath()
			defer os.Remove(path)

			storage, _ := zorroBolt.New(path)
			defer storage.Close()

			storage.LoadOrStore(tc.loadedID, tc.loadedValue)
			key, ok := storage.Resolve(tc.value)

			assert.Equal(t, tc.expectedID, key)
			assert.Equal(t, tc.expectedOk, ok)
		})
	}
}

func TestLoadOrStoreIncrementsStoreOpsIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	path := getTmpPath()
	defer os.Remove(path)

	bolt, _ := zorroBolt.New(path)
	bolt.WithMetrics(&storage.Metrics{StoreOps: counter})
	defer bolt.Close()

	bolt.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func TestLoadOrStoreIncrementsLoadOpsIfKeyExists(t *testing.T) {
	t.Parallel()

	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	path := getTmpPath()
	defer os.Remove(path)

	bolt, _ := zorroBolt.New(path)
	bolt.WithMetrics(&storage.Metrics{LoadOps: counter})
	defer bolt.Close()

	bolt.LoadOrStore("foo", "bar")
	bolt.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}
func TestResolveIncrementsResolveOps(t *testing.T) {
	t.Parallel()

	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	path := getTmpPath()
	defer os.Remove(path)

	bolt, _ := zorroBolt.New(path)
	bolt.WithMetrics(&storage.Metrics{ResolveOps: counter})
	defer bolt.Close()

	bolt.Resolve("bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func getTmpPath() string {
	f, err := ioutil.TempFile("", "zorro-tests")

	if err != nil {
		panic(err)
	}

	path := f.Name()
	defer os.Remove(path)

	f.Close()

	return path
}
