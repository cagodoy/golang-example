package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

func createPersonsHandler(ps PersonStorage) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// set headers
		w.Header().Set("Content-Type", "application/json")

		// // read body bytes
		// b, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	// 500? 
		// 	log.Fatalln("readAll error", err)
		// }

		// fmt.Println("BODY", string(b))
		
		// prepare struct for body
		body := &struct{
			Person *person.Person `json:"person"`
		}{}

		// // json.Marshal() // JSON.stringify
		// err = json.Unmarshal(b, &body) // JSON.Parse
		// if err != nil {
		// 	log.Fatalln("unmarshal error", err)
		// 	// error
		// }

		// parse body using json decoder
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			log.Fatalln("decoder.Decode( error", err)
			// error
		}

		// TODO(nb): validate json schema

		// created person in storage
		created, err := ps.Create(body.Person)
		if err != nil {
			log.Fatalln("ps.Create(body.Person) error", err)
		}

		// endpoint response
		res := response{
			Data: created,
		}

		// prepare response to api
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
