package personstorage

import (
	"encoding/json"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/darchlabs/api-example/pkg/storage"
	"github.com/teris-io/shortid"
)

type ps struct {
	DB *storage.DB
}

func New(s *storage.DB) *ps {
	return &ps{
		DB: s,
	}
}

func (s *ps) List() ([]*person.Person, error) {
	// // format the composed prefix key used in db
	// prefix := "Name:Age:"

	// data := make([]*Person, 0)

	// iter := s.storage.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	// for iter.Next() {
	// 	var pperson *Person

	// 	data = append(data, pperson)

	// }

	// iter.Release()
	// err := iter.Error()

	// if err != nil {
	// 	return nil, err
	// }

	// return data, nil
	return nil, nil
}

func (s *ps) Create(p *person.Person) (*person.Person, error) {
	// generate id for database
	id, err := shortid.Generate()
	if err != nil {
		return nil, err
	}

	// set generated id to person
	p.Id = id

	// JSON.stringify
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// save in database
	err = s.DB.DB.Put([]byte(id), b, nil)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// func (s *Storage) GetPersonByAge(age int64) ([]byte, error) {
// 	data, err := s.storage.DB.Get([]byte(fmt.Sprintf("%v", age)), nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, err

func (s *ps) Delete(id string) error {
	return nil
}
