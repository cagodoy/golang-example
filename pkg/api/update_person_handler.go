package api

import (
	"encoding/json"

	"github.com/darchlabs/api-example/pkg/person"
)

type UpdatePersonHandler struct {
	storage PersonStorage
}

func NewUpdatePersonHandler(ps PersonStorage) *UpdatePersonHandler {
	return &UpdatePersonHandler{storage: ps}
}

func (*UpdatePersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	id := ctx.r.FormValue("id")
	body := &struct {
		Person *person.Person `json:"person"`
	}{} // TODO(ca): What does this second '{}' means?

	decoder := json.NewDecoder(ctx.r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	// // Validate  req params
	// if id == "" {
	// 	log.Fatalln("No id provided")
	// }

	// // Validate body
	// if body.Person.Name == "" {
	// 	log.Fatalln("Name cannot be empty")
	// }

	// if body.Person.Age == 0 {
	// 	log.Fatalln("Age cannot be zero")
	// }

	pp, err := ctx.ps.UpdatePersonById(id, body.Person)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	return &handlerRes{pp, 200, nil}

}
