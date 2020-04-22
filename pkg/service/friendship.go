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

	// first check if the user one exists
	usrOne, err := repo.User().Get(model.User{ID: getBy.UserOneID}, f.db)
	if err != nil || usrOne.ID == "" {
		log.Printf("error getting the User with UserID %s:\n%s\n", getBy.UserOneID, err)
		return friendship, err
	}

	// then check if the user two exists
	usrTwo, err := repo.User().Get(model.User{ID: getBy.UserTwoID}, f.db)
	if err != nil || usrTwo.ID == "" {
		log.Printf("error getting the User with UserID %s:\n%s\n", getBy.UserTwoID, err)
		return friendship, err
	}

	// search the friendship between the users
	friendship, err = repo.Friendship().Get(getBy, f.db)
	if err != nil {
		log.Printf("error getting the Friendship by %+v:\n%s\n", getBy, err)
		return friendship, err
	}
	// TODO check if friendship is empty and return specific code

	return friendship, nil
}

func (f Friendship) Update(friendship model.Friendship) error {
	getBy := model.Friendship{
		UserOneID: friendship.UserOneID,
		UserTwoID: friendship.UserTwoID,
	}
	friendshipDB, err := repo.Friendship().Get(getBy, f.db)
	if err != nil {
		log.Printf("error getting the Friendship for UserOneID %s and UserTwoID %s:\n%s\n", friendship.UserOneID, friendship.UserTwoID, err)
		return err
	}
	if friendshipDB.UserOneID == "" || friendshipDB.UserTwoID == "" {
		// TODO return specific code
		return fmt.Errorf("error Friendship for UserOneID %s and UserTwoID %s not found", friendship.UserOneID, friendship.UserTwoID)
	}

	// Update status with the one received
	friendshipDB.Status = friendship.Status
	// UserTwoID received will be the user that answers the request (the receiver)
	friendshipDB.ActionUserID = friendship.UserTwoID
	if err := repo.Friendship().Update(friendshipDB, f.db); err != nil {
		log.Printf("error updating the Friendship %+v:\n%s\n", friendshipDB, err)
		return err
	}

	return nil
}
