package model

import "time"

type Group struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

func (Group) TableName() string {
	return "generator.group"
}

type UserGroup struct {
	UserID    string    `json:"userID"`
	GroupID   string    `json:"groupID"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

func (UserGroup) TableName() string {
	return "generator.user_group"
}

type Member struct {
	User
	IsAdmin bool `json:"isAdmin"`
}

type GroupComplete struct {
	Group
	Members []Member `json:"members"`
}
