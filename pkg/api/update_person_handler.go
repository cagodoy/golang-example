package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

func updatePersonHandler(ps PersonStorage) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.FormValue("id")
		body := &struct {
			Person *person.Person `json:"person"`
		}{} // TODO(ca): What does this second '{}' means?

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			log.Fatalln("decoder.Decode( error", err)
			// error
		}

		pp, err := ps.UpdatePersonById(id, body.Person)
		if err != nil {
			log.Fatalln("decoder.Decode error", err)
		}

		res := response{
			Data: pp,
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
