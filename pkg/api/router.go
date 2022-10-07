package api

import (
	"github.com/darchlabs/api-example/pkg/person"
	"github.com/darchlabs/api-example/pkg/storage"
	"github.com/julienschmidt/httprouter"
)

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func NewRouter(db *storage.S) *httprouter.Router {
	// initialize router
	router := httprouter.New()

	/* list persons */
	// Set a route for the func that lists persons in the s db
	router.GET("/api/v1/persons", listPersonsHandler(db))
	/* add persons */
	// Create person p var
	p := &person.Person{
		Name: "Nico",
		Age:  20,
	}
	// Set a route for the func that adds p persons to s db
	router.POST("/api/v1/persons", addPersonsHandler(s, p))

	return router
}
