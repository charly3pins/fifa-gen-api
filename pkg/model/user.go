package model

import "time"

type User struct {
	ID        string
	Name      string
	Active    bool
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}
