package main

import (
	"log"
	"net/http"

	"github.com/darchlabs/api-example/pkg/api"
	"github.com/darchlabs/api-example/pkg/storage"
)

func main() {

	// Open db
	db, err := storage.New("storage.db")
	if err != nil {
		log.Fatalf("Bad db opening %w", err) // fatal kills
	}

	// Initialize storage with leveldb db
	// s := person.New(db)

	// load router
	router := api.NewRouter(db)

	// run api server
	log.Println("runnning server...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}

	// storage.LevelDb()
}
