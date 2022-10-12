package api

import (
	"encoding/json"
	"log"
	"net/http"

	// personstorage "github.com/darchlabs/api-example/pkg/storage/person"

	"github.com/julienschmidt/httprouter"
)

func delPersonHandler(ps PersonStorage) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.FormValue("id")

		pp, err := ps.DeletePersonById(id)
		if err != nil {
			log.Fatalf("Error when trying to delete a person by id: %v", err)
		}

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
