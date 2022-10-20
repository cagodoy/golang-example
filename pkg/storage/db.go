package storage

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	DB *leveldb.DB
}

func New(filepath string) (*DB, error) {
	if filepath == "" {
		return nil, fmt.Errorf("%s", "Empty filepath param")
	}

	db, err := leveldb.OpenFile(fmt.Sprintf("./%s", filepath), nil)

	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}
