package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/darchlabs/api-example/pkg/api"
	"github.com/darchlabs/api-example/pkg/storage"
	personstorage "github.com/darchlabs/api-example/pkg/storage/person"
)

func main() {

	// load env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("invalid env PORT value")
	}
	
	// open db
	db, err := storage.New("storage.db")
	if err != nil {
		log.Fatalf("Bad db opening %s", err) // fatal kills
	}

	// intialize person storage
	s := personstorage.New(db)
	
	// load router
	router := api.NewRouter(s)

	// run api server
	log.Println("runnning server...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
