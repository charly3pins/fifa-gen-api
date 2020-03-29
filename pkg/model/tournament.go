package model

import "time"

type Tournament struct {
	ID             string
	Name           string
	Type           string
	NumPlayers     int
	NumTeamsPlayer int
	TimesAgainst   int `gorm:"default:NULL"`
	Round          int `gorm:"default:NULL"`
	GroupID        string
	CreatedAt      time.Time `gorm:"-"`
	UpdatedAt      time.Time `gorm:"-"`
}

type UserTournament struct {
	ID           string
	UserID       string
	TournamentID string
	TeamID       string
	CreatedAt    time.Time `gorm:"-"`
	UpdatedAt    time.Time `gorm:"-"`
}
