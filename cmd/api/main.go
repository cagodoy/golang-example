package main

import (
	"log"
	"net/http"

	"github.com/darchlabs/api-example/pkg/api"
)

func main() {
	// load router
	router := api.NewRouter()

	// run api server
	log.Println("runnning server...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}