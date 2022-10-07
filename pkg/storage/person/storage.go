package person

import (
	"fmt"

	"github.com/darchlabs/api-example/pkg/storage"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// Composition
type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage{
		storage: s,
	}
}

func (s Storage) GetPersons() ([]*Person, error) {
	// format the composed prefix key used in db
	prefix := "Name:Age:"

	data := make([]*Person, 0)

	iter := s.storage.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		var pperson *Person

		data = append(data, pperson)

	}

	iter.Release()
	err := iter.Error()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s Storage) NewPerson(p Person) (int64, error) {
	count := int64(0)

	keys := "Persons:Name:Age:"
	values := fmt.Sprintf(p.Name, p.Age)

	err := s.storage.DB.Put([]byte(keys), []byte(values), nil)

	if err != nil {
		return 0, err
	}

	// increase in one the counter
	count++

	return count, nil
}

func (s *Storage) GetPersonByAge(age int64) ([]byte, error) {
	data, err := s.storage.DB.Get([]byte(fmt.Sprintf("%v", age)), nil)

	if err != nil {
		return nil, err
	}

	return data, err
}
