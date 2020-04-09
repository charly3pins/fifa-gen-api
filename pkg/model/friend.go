package model

import "time"

type Friend struct {
	ID        string    `json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"-" gorm:"-"`
	UpdatedAt time.Time `json:"-" gorm:"-"`
}
