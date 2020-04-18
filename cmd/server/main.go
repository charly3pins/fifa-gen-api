package main

import (
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/handler"
	"github.com/charly3pins/fifa-gen-api/pkg/service"

	"github.com/gorilla/mux"
)

func main() {
	// Handlers
	fifaLeagueHandler := handler.NewFifaLeague(service.NewFifaLeague())
	fifaTeamHandler := handler.NewFifaTeam(service.NewFifaTeam())
	fifaPlayerHandler := handler.NewFifaPlayer(service.NewFifaPlayer())
	userHandler := handler.NewUser(service.NewUser(), service.NewFriendship())

	// Routes
	// FIFA information
	r := mux.NewRouter()
	r.HandleFunc("/fifa/leagues", fifaLeagueHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/teams", fifaTeamHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/teams/{id}", fifaTeamHandler.Get).Methods("GET")
	r.HandleFunc("/fifa/players", fifaPlayerHandler.Find).Methods("GET")
	r.HandleFunc("/fifa/players/{id}", fifaPlayerHandler.Get).Methods("GET")

	// Users
	// Login user (log in)
	r.HandleFunc("/token", userHandler.Login).Methods("POST") // TODO use JWT
	// Create user (sign up)
	r.HandleFunc("/users", userHandler.Create).Methods("POST")
	// Retrieve users (used in a search tool)
	r.HandleFunc("/users", userHandler.Find).Methods("GET")
	// Update user basic information (edit profile)
	r.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")

	// Friendships
	// Create a friend request for a user {id}
	r.HandleFunc("/users/{id}/friendships", userHandler.CreateFriendship).Methods("POST")
	// Find friendships for a user {id} [Query param filter={requested, pending, friends}]
	r.HandleFunc("/users/{id}/friendships", userHandler.FindFriendships).Methods("GET")
	// Answer a friendship request
	r.HandleFunc("/users/{id}/friendships", userHandler.UpdateFriendship).Methods("PUT")

	http.ListenAndServe(":8000", r)
}
