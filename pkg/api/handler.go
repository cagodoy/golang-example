package api

import (
	"encoding/json"
	"net/http"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

func listPersonsHandler(s *person.Storage) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// set headers
		w.Header().Set("Content-Type", "application/json")

		pp, err := s.GetPersons()
		if err != nil {
			return
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

func addPersonsHandler(s *person.Storage, p *person.Person) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// set headers
		w.Header().Set("Content-Type", "application/json")

		pp, err := s.NewPerson(*p)
		if err != nil {
			return
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
