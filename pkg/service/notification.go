package service

import (
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewNotification() Notification {
	// Database
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return Notification{
		db: db,
	}
}

type Notification struct {
	db *gorm.DB
}

func (n Notification) Find(userID string) (model.Notification, error) {
	var notifications model.Notification
	friends, err := repo.Friend().Find(userID, n.db)
	if err != nil {
		log.Printf("error finding the Friends for ID %s:\n%s\n", userID, err)
		return notifications, err
	}

	friendRequests := make([]model.User, len(friends))
	for k, f := range friends {
		usr, err := repo.User().Get(model.User{ID: f.Sender}, n.db)
		if err != nil {
			log.Printf("error getting User for ID %s:\n%s\n", f.Sender, err)
			return notifications, err
		}
		friendRequests[k] = usr
	}
	notifications.FriendRequests = friendRequests

	return notifications, nil
}
