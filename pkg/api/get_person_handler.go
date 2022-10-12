package api

import (
	"encoding/json"
	"net/http"

	// personstorage "github.com/darchlabs/api-example/pkg/storage/person"

	"github.com/julienschmidt/httprouter"
)

func getPersonHandler(ps PersonStorage) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// set headers
		w.Header().Set("Content-Type", "application/json")

		id := r.FormValue("id")

		pp, err := ps.GetPersonById(id)
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
