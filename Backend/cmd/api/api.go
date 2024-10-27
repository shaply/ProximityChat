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
	db   *mongo.Client // https://www.mongodb.com/resources/languages/golang?utm_source=google&utm_campaign=search_gs_pl_evergreen_atlas_language_prosp-nbnon_gic-null_amers-us_ps-all_dv-all_eng_lead&utm_term=golang%20database&utm_medium=cpc_paid_search&utm_ad=p&utm_ad_campaign_id=19248124983&adgroup=139647663730&cq_cmp=19248124983&gad_source=1&gclid=CjwKCAjwyfe4BhAWEiwAkIL8sP-3rb0dNlfKbdSowpD2fMYnbTPqeJiX3ae1xl8wh0wissEVgqxsjxoCgRAQAvD_BwE
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
