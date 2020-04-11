package service

import (
	"fmt"
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewFriend() Friend {
	// Database
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return Friend{
		db: db,
	}
}

type Friend struct {
	db *gorm.DB
}

func (f Friend) Create(friend model.Friend) (model.Friend, error) {
	getBy := model.Friend{
		Sender:   friend.Sender,
		Receiver: friend.Receiver,
	}
	friendDB, err := repo.Friend().Get(getBy, f.db)
	if err != nil {
		log.Printf("error getting the Friend for Sender %s and Receiver %s:\n%s\n", friend.Sender, friend.Receiver, err)
		return friend, err
	}
	if friendDB.ID != "" {
		// TODO return specific code
		return friend, fmt.Errorf("error duplicate Friend for Sender %s and Receiver %s", friend.Sender, friend.Receiver)
	}

	friendDB, err = repo.Friend().Create(friend, f.db)
	if err != nil {
		log.Printf("error creating the Friend %+v::\n%s\n", friend, err)
		return friend, err
	}

	return friendDB, nil
}

func (f Friend) Get(getBy model.Friend) (model.Friend, error) {
	friend, err := repo.Friend().Get(getBy, f.db)
	if err != nil {
		log.Printf("error getting the Friend %+v:\n%s\n", getBy, err)
		return friend, err
	}
	// TODO check if user exists if not return specific code

	return friend, nil
}
