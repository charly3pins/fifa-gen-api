package model

import "time"

type Group struct {
	ID        string
	Name      string
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

func (Group) TableName() string {
	return "generator.group"
}

type UserGroup struct {
	ID        string
	UserID    string
	GroupID   string
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

func (UserGroup) TableName() string {
	return "generator.user_group"
}
