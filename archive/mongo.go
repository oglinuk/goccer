package archive

import (
	"github.com/OGLinuk/xstore/mdbstore"
)

var (
	_ ArchiveStore = (*MongoStore)(nil)
)

type MongoStore struct {
	Instance *mdbstore.MDBStore
}

func NewMongoStore() *MongoStore {
	return &MongoStore{
		Instance: mdbstore.NewMDBStore(),
	}
}

func (ms *MongoStore) Archive() error {

	return nil
}
