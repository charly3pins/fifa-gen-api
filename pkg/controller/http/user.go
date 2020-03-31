package http

import (
	"encoding/json"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewUser(svc service.User) user {
	return user{svc: svc}
}

type user struct {
	svc service.User
}

func (u user) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err = u.svc.Create(user)
	if err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (u user) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/fifa/users/"):]

	var getBy model.User
	getBy.ID = id
	user, err := u.svc.Get(getBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = ""
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (u user) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var getBy model.User
	if err := json.NewDecoder(r.Body).Decode(&getBy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := u.svc.Get(getBy)
	if err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.ID == "" {
		// TODO check err code and return according message
		http.Error(w, "INVALID_LOGIN", http.StatusNotFound)
		return
	}
	user.Password = ""
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
