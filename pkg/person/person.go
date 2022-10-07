package person

import (
	"fmt"

	"github.com/darchlabs/api-example/pkg/storage"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// Composition
type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage{
		storage: s,
	}
}

func (s Storage) GetPersons() ([]string, error) {
	// format the composed prefix key used in db
	prefix := "Name:Age:"

	var pArray []string

	iter := s.storage.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		key := fmt.Sprintf(`%d`, iter.Key())
		value := fmt.Sprintf(`%d`, iter.Value())

		pArray = append(pArray, key, value)

	}

	iter.Release()
	err := iter.Error()

	if err != nil {
		return nil, err
	}

	return pArray, nil
}

func (s Storage) NewPerson(p Person) (int64, error) {
	count := int64(0)

	keys := "Name:Age:"
	values := fmt.Sprintf(p.Name, p.Age)

	err := s.storage.DB.Put([]byte(keys), []byte(values), nil)

	if err != nil {
		return 0, err
	}

	// increase in one the counter
	count++

	return count, nil
}

// func (p *Person) GetName() string {
// 	return p.Name
// }

// composition over inheritance
//
// type Teacher struct {
// 	Person Person `json:"person"`

// 	Subject string  `json:"subject"`
// }
