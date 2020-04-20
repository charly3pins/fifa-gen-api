package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/charly3pins/fifa-gen-api/pkg/model"
	"github.com/charly3pins/fifa-gen-api/pkg/service"
)

func NewUser(us service.User, fs service.Friendship) user {
	return user{
		userSvc:       us,
		friendshipSvc: fs,
	}
}

type user struct {
	userSvc       service.User
	friendshipSvc service.Friendship
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

	user, err = u.userSvc.Create(user)
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

func (u user) CreateFriendship(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friendship, err := u.friendshipSvc.Create(friendship)
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
func (u user) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var findBy model.User
	keys, ok := r.URL.Query()["username"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'username' is missing")
		return
	}
	findBy.Username = keys[0]
	var usrs []model.User
	usrs, err := u.userSvc.Find(findBy) // TODO check how to do this findBy optional
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(usrs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (u user) FindFriendships(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	userID := params["id"]

	filter := ""
	keys, ok := r.URL.Query()["filter"]
	if ok {
		filter = keys[0]
	}

	users, err := u.friendshipSvc.Find(userID, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (u user) GetFriendship(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	userID := params["id"]
	otherUserID := params["otherUserID"]

	getBy := model.Friendship{
		UserOneID: userID,
		UserTwoID: otherUserID,
	}
	friendship, err := u.friendshipSvc.Get(getBy)
	if err != nil {
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
	user, err := u.userSvc.Get(getBy)
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

func (u user) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.userSvc.Update(user); err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte{})
}

func (u user) UpdateFriendship(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.friendshipSvc.Update(friendship); err != nil {
		// TODO check err code and return according message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte{})
}
