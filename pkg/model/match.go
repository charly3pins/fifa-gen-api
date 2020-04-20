package model

import "time"

type Match struct {
	ID        string
	FixtureID string
	Home      string
	Away      string
	Played    bool
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

func (Match) TableName() string {
	return "generator.match"
}
