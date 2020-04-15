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
		return friendship, fmt.Errorf("error duplicate Friend for UserOneID %s and UserTwoID %s", friendship.UserOneID, friendship.UserTwoID)
	}

	friendshipDB, err = repo.Friendship().Create(friendship, f.db)
	if err != nil {
		log.Printf("error creating the Friend %+v:\n%s\n", friendship, err)
		return friendship, err
	}

	return friendshipDB, nil
}

func (f Friendship) Get(getBy model.Friendship) (model.Friendship, error) {
	friendship, err := repo.Friendship().Get(getBy, f.db)
	if err != nil {
		log.Printf("error getting the Friendship %+v:\n%s\n", getBy, err)
		return friendship, err
	}
	// TODO check if user exists if not return specific code

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
