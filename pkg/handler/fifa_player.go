package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewFifaPlayer(svc service.FifaPlayer) fifaPlayer {
	return fifaPlayer{svc: svc}
}

type fifaPlayer struct {
	svc service.FifaPlayer
}

func (fl fifaPlayer) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/fifa/players/"):]

	var getBy model.FifaPlayer
	getBy.ID = id
	f, err := fl.svc.Get(getBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (fl fifaPlayer) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var findBy model.FifaPlayer
	keys, ok := r.URL.Query()["teamID"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'teamID' is missing")
		return
	}
	findBy.TeamID = keys[0]

	var f []model.FifaPlayer
	f, err := fl.svc.Find(findBy) // TODO check how to do this findBy optional
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
