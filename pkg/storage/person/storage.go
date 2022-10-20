package personstorage

import (
	"encoding/json"
	"fmt"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/darchlabs/api-example/pkg/storage"
	"github.com/teris-io/shortid"
)

type ps struct {
	DB *storage.DB
	id *shortid.Shortid
}

// Initialize an instance of DB (leveldv in this case)
func New(s *storage.DB) (*ps, error) {
	idGenerator, err := shortid.New(1, "-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_", 2342)
	if err != nil {
		return nil, err
	}

	return &ps{
		DB: s,
		id: idGenerator,
	}, nil
}

// Functions for interacting with the storage
func (s *ps) List() ([]*person.Person, error) {
	// Create slice where the values will be returned
	data := make([]*person.Person, 0)

	// Iterate over the values and append them to the slice
	iter := s.DB.DB.NewIterator(nil, nil)
	for iter.Next() {
		var p *person.Person
		err := json.Unmarshal(iter.Value(), &p)
		if err != nil {
			return nil, err
		}

		data = append(data, p)
	}
	iter.Release()

	err := iter.Error()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ps) Create(p *person.Person) (*person.Person, error) {
	if p.Name == "" {
		return nil, fmt.Errorf("%s", "Empty Name param in person")
	}

	if p.Age == 0 {
		return nil, fmt.Errorf("%s", "Age param in person cannot be zero")
	}

	// TODO(nb): Add validation if the person already exists, it shouldn't be created

	// generate id for database
	id, err := s.id.Generate()
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

func (s *ps) GetPersonById(id string) (*person.Person, error) {
	if id == "" {
		return nil, fmt.Errorf("%s", "Empty filepath param")
	}

	data, err := s.DB.DB.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}

	var p *person.Person

	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (s *ps) UpdatePersonById(id string, p *person.Person) (*person.Person, error) {
	if id == "" {
		return nil, fmt.Errorf("%s", "Empty filepath param")
	}

	if p.Name == "" {
		return nil, fmt.Errorf("%s", "Empty Name param in person")
	}

	if p.Age == 0 {
		return nil, fmt.Errorf("%s", "Age param in person cannot be zero")
	}

	p.Id = id
	// If a person with that id doesn't exist, it should fail
	_, err := s.GetPersonById(p.Id)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	err = s.DB.DB.Put([]byte(id), b, nil)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *ps) DeletePersonById(id string) (*person.Person, error) {
	if id == "" {
		return nil, fmt.Errorf("%s", "Empty filepath param")
	}

	data, err := s.DB.DB.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}

	var p *person.Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	err = s.DB.DB.Delete([]byte(id), nil)
	return p, err
}
