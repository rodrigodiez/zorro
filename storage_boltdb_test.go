package zorro

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
)

func TestNewBoltReturnsErrIfCantOpen(t *testing.T) {
	storage, err := NewBoltDBStorage("/a/path/that/does/not/exist")

	assert.Nil(t, storage)
	assert.NotNil(t, err)
}

func TestCloseClosesTheDB(t *testing.T) {
	t.Skip("Not sure how to do this yet")
}

func TestNewBoltCreatesIdsAndMAsksBuckets(t *testing.T) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	storage.Close()

	db, _ := bolt.Open(path, 0600, nil)
	db.View(func(tx *bolt.Tx) error {
		bIds := tx.Bucket([]byte("ids"))
		bMasks := tx.Bucket([]byte("masks"))

		assert.NotNil(t, bIds, "ids bucket does not exist")
		assert.NotNil(t, bMasks, "masks bucket does not exist")

		return nil
	})

	db.Close()
}

func TestBoltLoadOrStoreReTestturnsMaskAndFalseIfIdDoesNotExist(t *testing.T) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	defer storage.Close()

	mask, loaded := storage.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", mask)
	assert.Equal(t, false, loaded)
}

func TestBoltLoadOrStoreReturnsActualMaskAndTrueIfKeyExists(t *testing.T) {
	path := getTmpPath()
	defer os.Remove(path)

	storage, _ := NewBoltDBStorage(path)
	defer storage.Close()

	storage.LoadOrStore("foo", "bar")
	mask, loaded := storage.LoadOrStore("foo", "baz")

	assert.Equal(t, "bar", mask)
	assert.Equal(t, true, loaded)
}

func TestBoltResolve(t *testing.T) {
	tt := []struct {
		name       string
		loadedID   string
		loadedMask string
		mask       string
		expectedID string
		expectedOk bool
	}{
		{name: "Id exists", loadedID: "foo", loadedMask: "bar", mask: "bar", expectedID: "foo", expectedOk: true},
		{name: "Id does not exist", loadedID: "foo", loadedMask: "bar", mask: "baz", expectedID: "", expectedOk: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			path := getTmpPath()
			defer os.Remove(path)

			storage, _ := NewBoltDBStorage(path)
			defer storage.Close()

			storage.LoadOrStore(tc.loadedID, tc.loadedMask)
			id, ok := storage.Resolve(tc.mask)

			assert.Equal(t, tc.expectedID, id)
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
