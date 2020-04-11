package http

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

	var friend model.Friendship
	err := json.NewDecoder(r.Body).Decode(&friend)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friend, err = f.svc.Create(friend)
	if err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(friend)
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
	log.Println("getBy", getBy)
	friendship, err := f.svc.Get(getBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("friendship", friendship)
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
