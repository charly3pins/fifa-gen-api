package main

import (
	"net/http"

	handler "github.com/charly3pins/fifa-gen-api/pkg/controller/http"
	"github.com/charly3pins/fifa-gen-api/pkg/service"

	"github.com/gorilla/mux"
)

func main() {
	// Handlers
	fifaLeagueHandler := handler.NewFifaLeague(service.NewFifaLeague())
	fifaTeamHandler := handler.NewFifaTeam(service.NewFifaTeam())
	fifaPlayerHandler := handler.NewFifaPlayer(service.NewFifaPlayer())
	userHandler := handler.NewUser(service.NewUser())
	groupHandler := handler.NewGroup(service.NewGroup())

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/fifa/leagues", fifaLeagueHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/teams", fifaTeamHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/teams/{id}", fifaTeamHandler.Get).Methods("GET")
	r.HandleFunc("/fifa/players", fifaPlayerHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/players/{id}", fifaPlayerHandler.Get).Methods("GET")

	r.HandleFunc("/users", userHandler.Create).Methods("POST")
	r.HandleFunc("/token", userHandler.Login).Methods("POST") // TODO use jwt

	r.HandleFunc("/groups", groupHandler.Find).Methods("GET")
	r.HandleFunc("/groups/{id}", groupHandler.Get).Methods("GET")


	http.ListenAndServe(":8000", r)
}
