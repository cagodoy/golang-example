package api

import (
	"github.com/darchlabs/api-example/pkg/person"
	"github.com/julienschmidt/httprouter"
)

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type PersonStorage interface {
	List() ([]*person.Person, error)
	Create(p *person.Person) (*person.Person, error)
}

func NewRouter(ps PersonStorage) *httprouter.Router {
	// initialize router
	router := httprouter.New()

	// Set a route for the func that lists persons in the s storage
	router.GET("/api/v1/persons", listPersonsHandler(ps))

	// Set a route for the func that adds p persons to s db
	router.POST("/api/v1/persons", createPersonsHandler(ps))

	return router
}
