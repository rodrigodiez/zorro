package boltdb_test

import (
	"testing"

	"github.com/boltdb/bolt"

	zorroBoltdb "github.com/rodrigodiez/zorro/pkg/storage/boltdb"
)

func TestNewClientAdapterImplementsClientAdapter(t *testing.T) {
	t.Parallel()

	var _ zorroBoltdb.ClientAdapter = zorroBoltdb.NewClientAdapter(&bolt.DB{})
}
