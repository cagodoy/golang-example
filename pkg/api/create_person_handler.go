package api

import (
	"encoding/json"
	"fmt"

	"github.com/darchlabs/api-example/pkg/person"
)

type CreatePersonHandler struct {
	storage PersonStorage
}

func NewCreatePersonHandler(ps PersonStorage) *CreatePersonHandler {
	return &CreatePersonHandler{
		storage: ps,
	}
}

func (cp CreatePersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	// set headers
	ctx.w.Header().Set("Content-Type", "application/json")

	// prepare struct for body
	body := &struct {
		Person *person.Person `json:"person"`
	}{}

	// parse body using json decoder
	decoder := json.NewDecoder(ctx.r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	// Validate json schema data
	personBody := body.Person
	if personBody == nil {
		return &handlerRes{err.Error(), 500, err}
	}

	name := body.Person.Name
	nameType := fmt.Sprintf("%T", name)
	if name == "" || nameType != "string" {
		return &handlerRes{"Invalid name param", 500, err}
	}

	age := body.Person.Age
	ageType := fmt.Sprintf("%T", body.Person.Age)
	if age == 0 || ageType != "int64" {
		return &handlerRes{"Invalid age param", 500, err}
	}

	// created person in storage
	created, err := cp.storage.Create(body.Person)
	if err != nil {
		return &handlerRes{"Error creating person", 500, err}
	}

	return &handlerRes{created, 200, nil}
}
