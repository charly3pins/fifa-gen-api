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
	pendingReqs, err := repo.Friendship().PendingRequests(userID, n.db)
	if err != nil {
		log.Printf("error finding the Friends for ID %s:\n%s\n", userID, err)
		return notifications, err
	}

	friendRequests := make([]model.User, len(pendingReqs))
	for k, pr := range pendingReqs {
		var getBy model.User
		// Search the users of the friend requests that are not the user
		if pr.UserOneID == userID {
			getBy.ID = pr.UserTwoID
		} else {
			getBy.ID = pr.UserOneID
		}
		usr, err := repo.User().Get(getBy, n.db)
		if err != nil {
			log.Printf("error getting User for ID %s:\n%s\n", getBy.ID, err)
			return notifications, err
		}
		friendRequests[k] = usr
	}
	notifications.FriendRequests = friendRequests

	return notifications, nil
}
