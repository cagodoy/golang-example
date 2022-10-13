package api

import (
	"encoding/json"
	"net/http"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

// Define response structure
type response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// Define handler's response
type handlerRes struct {
	Payload    string
	HttpStatus int
	err        error
}

// Define handler's required context
type handlerCtx struct {
	ps PersonStorage
	w  http.ResponseWriter
	r  http.Request
}

type handler func(c *handlerCtx) handlerRes

// Define ps methods
type PersonStorage interface {
	Create(p *person.Person) (*person.Person, error)
	List() ([]*person.Person, error)
	GetPersonById(id string) (*person.Person, error)
	UpdatePersonById(id string, p *person.Person) (*person.Person, error)
	DeletePersonById(id string) (*person.Person, error)
}

// Router for managing the routes to the handlers, it receives the instance of person storage (it must have the interface methods)
func NewRouter(ps PersonStorage) *httprouter.Router {
	// initialize router
	router := httprouter.New()

	/// HERE IS THE ERROR: I can't pass the func as parameter for the decorator
	// Set a route for the func that get p persons from s db by the id
	// router.GET("/api/v1/person", HandleFunc(getPersonHandler, &ps))

	return router
}

func HandleFunc(fn handler, ps *PersonStorage) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := handlerCtx{ps: *ps, w: w, r: *r}
		handlerRes := fn(&ctx)
		payload, statusCode, err := handlerRes.Payload, handlerRes.HttpStatus, handlerRes.err

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := response{
			Data: payload,
			Meta: statusCode,
		}

		json.NewEncoder(w).Encode(res)
	}
}

// Set a route for the func that lists persons in the s storage
// router.GET("/api/v1/persons", HandleFunc(listPersonsHandler()))

// Set a route for the func that adds p persons to s db
// router.PUT("/api/v1/persons", updatePersonHandler)

// // Set a route for the func that adds p persons to s db
// router.POST("/api/v1/persons", createPersonsHandler)

// // Set a route for delete a person from s db by the id
// router.DELETE("/api/v1/person", delPersonHandler)
