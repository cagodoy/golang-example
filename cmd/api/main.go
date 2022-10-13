package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/darchlabs/api-example/pkg/api"
	"github.com/darchlabs/api-example/pkg/storage"
	personstorage "github.com/darchlabs/api-example/pkg/storage/person"
	"github.com/joho/godotenv"
)

func main() {

	// load env values
	godotenv.Load(".env")

	// get env values
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("invalid env PORT value")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatalf("invalid db path")
	}

	// open db
	db, err := storage.New(dbPath)
	if err != nil {
		log.Fatalf("Bad db opening %s", err)
	}

	// intialize person storage
	s := personstorage.New(db)

	// load router
	router := api.NewRouter(s)

	// run api server
	log.Println("runnning server ...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
