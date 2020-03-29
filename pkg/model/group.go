package model

import "time"

type Group struct {
	ID        string
	Name      string
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

type UserGroup struct {
	ID        string
	UserID    string
	GroupID   string
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}
