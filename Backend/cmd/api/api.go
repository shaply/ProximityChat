package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" // Need to run go get github.com/gorilla/mux to install this package
	"github.com/shaply/ProximityChat/Backend/service/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr string
	db   *mongo.Client
}

func NewAPIServer(addr string, db *mongo.Client) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter() // Creates new router for API server
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter) // Handles the routes of the router

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

// https://www.youtube.com/watch?v=7VLmLOiQ3ck
// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
