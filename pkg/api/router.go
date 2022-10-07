package api

import (
	"fmt"

	"github.com/darchlabs/api-example/pkg/person"
	"github.com/darchlabs/api-example/pkg/storage"
	"github.com/julienschmidt/httprouter"
)

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func NewRouter() *httprouter.Router {
	// initialize router
	router := httprouter.New()

	// Open db
	db, err := storage.New("storage.db")

	if err != nil {
		fmt.Println("Bad db opening %w", err)
		return nil
	}

	// Initialize storage with leveldb db
	s := person.New(db)

	/* list persons */
	// Set a route for the func that lists persons in the s db
	router.GET("/api/v1/persons", listPersonsHandler(s))
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
