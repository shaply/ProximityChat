package main

import (
	"log"

	"github.com/shaply/ProximityChat/Backend/cmd/api"
)

func main() {
	server := api.NewAPIServer("localhost:8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
