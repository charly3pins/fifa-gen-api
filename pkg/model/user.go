package model

import "time"

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Active         bool      `json:"active"`
	ProfilePicture string    `json:"profilePicture"`
	CreatedAt      time.Time `json:"-" gorm:"-"`
	UpdatedAt      time.Time `json:"-" gorm:"-"`
}

func (User) TableName() string {
	return "generator.user"
}

const (
	StatusCodeRequested = 0
	StatusCodeAccepted  = 1
	StatusCodeDeclined  = 2
	StatusCodeBlocked   = 3

	FilterRequested = "requested"
	FilterPending   = "pending"
	FilterFriends   = "friends"
)

type Friendship struct {
	UserOneID    string    `json:"userOneID"`
	UserTwoID    string    `json:"userTwoID"`
	Status       int       `json:"status"`
	ActionUserID string    `json:"actionUserID"`
	CreatedAt    time.Time `json:"-" gorm:"-"`
	UpdatedAt    time.Time `json:"-" gorm:"-"`
}

func (Friendship) TableName() string {
	return "generator.friendship"
}
