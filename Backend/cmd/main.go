package main

import (
	"log"

	"github.com/shaply/ProximityChat/Backend/cmd/api"
	"github.com/shaply/ProximityChat/Backend/db"
)

func main() {
	// Initialize the MongoDB storage connection
	database, err := db.InitiateConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the API server
	server := api.NewAPIServer("localhost:8080", database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
