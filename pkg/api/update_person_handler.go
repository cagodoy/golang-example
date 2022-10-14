package api

import (
	"encoding/json"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/go-playground/validator/v10"
)

type UpdatePersonHandlerInputs struct {
	Name string `validate:"required"`
	Age  int64  `validate:"required"`
}

type UpdatePersonHandler struct {
	storage PersonStorage
}

func NewUpdatePersonHandler(ps PersonStorage) *UpdatePersonHandler {
	return &UpdatePersonHandler{storage: ps}
}

func (*UpdatePersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	//Get and parse id req param
	id := ctx.r.FormValue("id")

	// Validate body
	validate := validator.New()
	validate.Struct(ctx.ps)

	// Define structure for parse body
	body := &struct {
		Person *person.Person `json:"person"`
	}{} // TODO(ca): What does this second '{}' means?

	// Parse body
	decoder := json.NewDecoder(ctx.r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	// Update person in storage and get it as return
	pp, err := ctx.ps.UpdatePersonById(id, body.Person)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	return &handlerRes{pp, 200, nil}
}
