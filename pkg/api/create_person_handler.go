package api

import (
	"encoding/json"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/go-playground/validator/v10"
)

type CreatePersonHandlerInputs struct {
	Name string `validate:"required"`
	Age  int64  `validate:"required"`
}

type CreatePersonHandler struct {
	storage PersonStorage
}

func NewCreatePersonHandler(ps PersonStorage) *CreatePersonHandler {
	return &CreatePersonHandler{
		storage: ps,
	}
}

func (cp CreatePersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	// Validate json schema data
	validate := validator.New()
	validate.Struct(ctx.ps)

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

	// created person in storage
	created, err := cp.storage.Create(body.Person)
	if err != nil {
		return &handlerRes{"Error creating person", 500, err}
	}

	return &handlerRes{created, 200, nil}
}
