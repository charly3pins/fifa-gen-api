package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewFifaTeam(svc service.FifaTeam) fifaTeam {
	return fifaTeam{svc: svc}
}

type fifaTeam struct {
	svc service.FifaTeam
}

func (fl fifaTeam) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/fifa/teams/"):]

	var getBy model.FifaTeam
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

func (fl fifaTeam) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var findBy model.FifaTeam
	keys, ok := r.URL.Query()["leagueID"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'leagueID' is missing")
		return
	}
	findBy.LeagueID = keys[0]

	var f []model.FifaTeam
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
