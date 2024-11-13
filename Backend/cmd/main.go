package main

import (
	"fmt"
	"log"

	"github.com/shaply/ProximityChat/Backend/cmd/api"
	"github.com/shaply/ProximityChat/Backend/config"
	"github.com/shaply/ProximityChat/Backend/db"
)

func main() {
	// Initialize the MongoDB storage connection
	database, err := db.InitiateConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the API server
	server := api.NewAPIServer(fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port), database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
