package api

import (
	"encoding/json"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/go-playground/validator/v10"
)

type CreatePersonHandler struct {
	storage PersonStorage
	// Person  *person.Person `json:"person"`
}

func NewCreatePersonHandler(ps PersonStorage) *CreatePersonHandler {
	return &CreatePersonHandler{
		storage: ps,
		// Person:  &person.Person{},
	}
}

func (cp CreatePersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	// Validate json schema data
	validate := validator.New()
	validate.Struct(ctx.ps)

	// // Define structure for parse body
	// body := &person.Person{}

	// Define structure for parse body
	body := &struct {
		Person *person.Person `json:"person"`
	}{}

	// Parse body
	decoder := json.NewDecoder(ctx.r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	// create a person in storage
	created, err := ctx.ps.Create(body.Person)
	if err != nil {
		return &handlerRes{"Error creating person", 500, err}
	}

	return &handlerRes{created, 200, nil}
}
