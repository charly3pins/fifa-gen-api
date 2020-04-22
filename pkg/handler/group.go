package handler

import (
	"encoding/json"
	"net/http"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"

	"github.com/gorilla/mux"
)

func NewGroup(gs service.Group, us service.User) group {
	return group{
		groupSvc: gs,
		userSvc:  us,
	}
}

type group struct {
	groupSvc service.Group
	userSvc  service.User
}

func (g group) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var groupComplete model.GroupComplete
	if err := json.NewDecoder(r.Body).Decode(&groupComplete); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group, err := g.groupSvc.Create(groupComplete)
	if err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (g group) Find(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 	return
	// }

	// var findBy model.User
	// keys, ok := r.URL.Query()["username"]
	// if !ok || len(keys[0]) < 1 {
	// 	log.Println("Url Param 'username' is missing")
	// 	return
	// }
	// findBy.Username = keys[0]
	// var usrs []model.User
	// usrs, err := u.userSvc.Find(findBy) // TODO check how to do this findBy optional
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// b, err := json.Marshal(usrs)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(b)
}

func (g group) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	getBy := model.Group{
		ID: id,
	}
	group, err := g.groupSvc.Get(getBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (g group) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var groupComplete model.GroupComplete
	if err := json.NewDecoder(r.Body).Decode(&groupComplete); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := g.groupSvc.Update(groupComplete); err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte{})
}
