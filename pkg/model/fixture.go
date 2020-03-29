package model

import "time"

type Fixture struct {
	ID           string
	Name         string
	TournamentID string
	CreatedAt    time.Time `gorm:"-"`
	UpdatedAt    time.Time `gorm:"-"`
}
