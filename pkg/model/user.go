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
