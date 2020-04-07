package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewGroup(svc service.Group) group {
	return group{svc: svc}
}

type group struct {
	svc service.Group
}

func (g group) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/groups/"):]

	var getBy model.Group
	getBy.ID = id
	f, err := g.svc.Get(getBy)
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

func (gr group) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var findBy model.Group
	keys, ok := r.URL.Query()["ID"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'ID' is missing")
		return
	}
	findBy.ID = keys[0]

	var group []model.Group
	group, err := gr.svc.Find(findBy) // TODO check how to do this findBy optional
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	g, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(g)
}

func (gr group) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var group model.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group, err = gr.svc.Create(group)
	if err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	g, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(g)
}
