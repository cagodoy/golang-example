package api

import (
	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

// Define response structure
type response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

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

	// Set a route for the func that lists persons in the s storage
	router.GET("/api/v1/persons", listPersonsHandler(ps))

	// Set a route for the func that get p persons from s db by the id
	router.GET("/api/v1/person", getPersonHandler(ps))

	// Set a route for the func that adds p persons to s db
	router.PUT("/api/v1/persons", updatePersonHandler(ps))

	// Set a route for the func that adds p persons to s db
	router.POST("/api/v1/persons", createPersonsHandler(ps))

	// Set a route for delete a person from s db by the id
	router.DELETE("/api/v1/person", delPersonHandler(ps))

	return router
}
