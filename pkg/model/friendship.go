package model

import "time"

const (
	StatusCodePending  = 0
	StatusCodeAccepted = 1
	StatusCodeDeclined = 2
	StatusCodeBlocked  = 3
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
