package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewFriend(svc service.Friend) friend {
	return friend{svc: svc}
}

type friend struct {
	svc service.Friend
}

func (f friend) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var friend model.Friend
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

func (f friend) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// TODO check how to improve the query param extraction
	var getBy model.Friend
	keys, ok := r.URL.Query()["sender"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'sender' is missing")
		return
	}
	getBy.Sender = keys[0]

	keys, ok = r.URL.Query()["receiver"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'receiver' is missing")
		return
	}
	getBy.Receiver = keys[0]

	friend, err := f.svc.Get(getBy)
	if err != nil {
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

func (f friend) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var findBy model.Friend
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}
	findBy.Receiver = keys[0]

	friends, err := f.svc.Find(findBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(len(friends))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
