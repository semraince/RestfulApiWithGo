package inmemory

import (
	"github.com/hashicorp/go-memdb"
)

type KeyValue struct {
	Key   string
	Value string
}

var inMemDb *memdb.MemDB

func Init() error {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"keyvalue": &memdb.TableSchema{
				Name: "keyvalue",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Key"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	inMemDb = db
	return nil
}

func Write(pair *KeyValue) error {
	transact := inMemDb.Txn(true)
	if err := transact.Insert("keyvalue", pair); err != nil {
		transact.Abort()
		return err
	}
	transact.Commit()
	return nil
}

func Read(key string) (*KeyValue, error) {
	transact := inMemDb.Txn(false)
	defer transact.Abort()

	itrx, err := transact.First("keyvalue", "id", key)
	if err != nil {
		return nil, err
	} else if itrx == nil {
		return nil, nil
	}
	return itrx.(*KeyValue), nil
}
