package api

import (
	"encoding/json"
	"net/http"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

func listPersonsHandler() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// set headers
		w.Header().Set("Content-Type", "application/json")

		// list persons
		// TODO(ca): get persons from database
		// pp := make([]*person.Person, 0)
		// pp = append(pp, person.NewPerson("nico", 20))
		// pp = append(pp, person.NewPerson("camilo", 33))

		pp := []*person.Person{
			person.NewPerson("nico", 20),
			person.NewPerson("camilo", 33),
		}

		// get values to response
		res := response{
			Data: pp,
		}

		// prepare response to api
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}