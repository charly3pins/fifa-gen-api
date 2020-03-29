package model

import "time"

type Goal struct {
	ID        string
	MatchID   string
	PlayerID  string
	Type      string    `gorm:"default:NULL"`
	Minute    int       `gorm:"default:NULL"`
	CreatedAt time.Time `gorm:"-"`
}
