package storage

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	DB *leveldb.DB
}

func New(filepath string) (*DB, error) {
	db, err := leveldb.OpenFile(fmt.Sprintf("./%s", filepath), nil)

	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}
