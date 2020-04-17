package service

import (
	"fmt"
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewFriendship() Friendship {
	// Database
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return Friendship{
		db: db,
	}
}

type Friendship struct {
	db *gorm.DB
}

func (f Friendship) Create(friendship model.Friendship) (model.Friendship, error) {
	// UserOneID will be the user that sends the request
	// UserTwoID will be the user that receives the request
	getBy := model.Friendship{
		UserOneID: friendship.UserOneID,
		UserTwoID: friendship.UserTwoID,
	}
	friendshipDB, err := repo.Friendship().Get(getBy, f.db)
	if err != nil {
		log.Printf("error getting the Friendship for UserOneID %s and UserTwoID %s:\n%s\n", friendship.UserOneID, friendship.UserTwoID, err)
		return friendship, err
	}
	if friendshipDB.UserOneID != "" || friendshipDB.UserTwoID != "" {
		// TODO return specific code
		return friendship, fmt.Errorf("error duplicate Friendship for UserOneID %s and UserTwoID %s", friendship.UserOneID, friendship.UserTwoID)
	}

	friendshipDB, err = repo.Friendship().Create(friendship, f.db)
	if err != nil {
		log.Printf("error creating the Friendship %+v:\n%s\n", friendship, err)
		return friendship, err
	}

	return friendshipDB, nil
}

func (f Friendship) Find(userID, filter string) ([]model.User, error) {
	var users []model.User

	// first check if the users exists
	usr, err := repo.User().Get(model.User{ID: userID}, f.db)
	if err != nil || usr.ID == "" {
		log.Printf("error getting the User with UserID %s:\n%s\n", userID, err)
		return users, err
	}

	// Find friendship for that user and the filter selected
	users, err = repo.Friendship().Find(userID, filter, f.db)
	if err != nil {
		log.Printf("error getting the Friendship by UserID%s and filter %s :\n%s\n", userID, filter, err)
		return users, err
	}

	return users, nil
}

func (f Friendship) Get(getBy model.Friendship) (model.Friendship, error) {
	var friendship model.Friendship
	return friendship, nil
}

func (f Friendship) Update(friendship model.Friendship) error {
	return nil
}
