package http

import (
	"encoding/json"
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

	id := r.URL.Path[len("/friends/"):]

	var getBy model.Friend
	getBy.ID = id
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
