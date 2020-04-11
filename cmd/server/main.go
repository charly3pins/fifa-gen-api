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
	friendshipHandler := handler.NewFriendship(service.NewFriendship())
	notificationHandler := handler.NewNotification(service.NewNotification())

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/fifa/leagues", fifaLeagueHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/teams", fifaTeamHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/teams/{id}", fifaTeamHandler.Get).Methods("GET")
	r.HandleFunc("/fifa/players", fifaPlayerHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/players/{id}", fifaPlayerHandler.Get).Methods("GET")

	r.HandleFunc("/token", userHandler.Login).Methods("POST") // TODO use jwt
	r.HandleFunc("/users", userHandler.Create).Methods("POST")
	r.HandleFunc("/users", userHandler.Find).Methods("GET")

	r.HandleFunc("/friendship", friendshipHandler.Create).Methods("POST")
	r.HandleFunc("/friendship", friendshipHandler.Get).Methods("GET")

	// TODO check how to improve this method
	r.HandleFunc("/notifications", notificationHandler.Find).Methods("GET")

	http.ListenAndServe(":8000", r)
}
