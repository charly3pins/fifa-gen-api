package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewFriendship(svc service.Friendship) friendship {
	return friendship{svc: svc}
}

type friendship struct {
	svc service.Friendship
}

func (f friendship) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friendship, err := f.svc.Create(friendship)
	if err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(friendship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (f friendship) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// TODO check how to improve the query param extraction
	var getBy model.Friendship
	keys, ok := r.URL.Query()["userOne"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'userOne' is missing")
		return
	}
	getBy.UserOneID = keys[0]

	keys, ok = r.URL.Query()["userTwo"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'userTwo' is missing")
		return
	}
	getBy.UserTwoID = keys[0]

	friendship, err := f.svc.Get(getBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if friendship.UserOneID == "" {
		w.Write([]byte{})
		return
	}
	b, err := json.Marshal(friendship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (f friendship) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := f.svc.Update(friendship); err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte{})
}
