package zorro

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
)

func TestNewBoltImplementsStorageCloser(t *testing.T) {
	var storage StorageCloser

	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ = NewBoltDBStorage(path)
	storage.Close()
}

func TestNewBoltReturnsErrIfCantOpen(t *testing.T) {
	t.Parallel()

	storage, err := NewBoltDBStorage("/a/path/that/does/not/exist")

	assert.Nil(t, storage)
	assert.NotNil(t, err)
}

func TestCloseClosesTheDB(t *testing.T) {
	t.Skip("Not sure how to do this yet")
}

func TestNewBoltCreatesKeysAndValuesBuckets(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
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

func TestBoltLoadOrStoreReTestturnsValueAndFalseIfIdDoesNotExist(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	defer storage.Close()

	value, loaded := storage.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Equal(t, false, loaded)
}

func TestBoltLoadOrStoreReturnsActualValueAndTrueIfKeyExists(t *testing.T) {
	t.Parallel()

	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	defer storage.Close()

	storage.LoadOrStore("foo", "bar")
	value, loaded := storage.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", value)
	assert.Equal(t, true, loaded)
}

func TestBoltResolve(t *testing.T) {
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

			storage, _ := NewBoltDBStorage(path)
			defer storage.Close()

			storage.LoadOrStore(tc.loadedID, tc.loadedValue)
			key, ok := storage.Resolve(tc.value)

			assert.Equal(t, tc.expectedID, key)
			assert.Equal(t, tc.expectedOk, ok)
		})
	}
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
