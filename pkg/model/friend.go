package model

import "time"

const (
	AcceptedState  = "ACCEPTED"
	RequestedState = "REQUESTED"
	PendingState   = "PENDING"
)

type Friend struct {
	ID        string    `json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"-" gorm:"-"`
	UpdatedAt time.Time `json:"-" gorm:"-"`
}

func (Friend) TableName() string {
	return "generator.friend"
}
