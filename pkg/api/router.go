package api

import (
	"github.com/julienschmidt/httprouter"
)

type response struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func NewRouter() *httprouter.Router {
	// initialize router
	router := httprouter.New()
	
	// list persons
	router.GET("/api/v1/persons", listPersonsHandler()) 

	return router
}