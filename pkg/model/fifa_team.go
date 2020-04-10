package model

type FifaTeam struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShieldSrc string `json:"shieldSrc" gorm:"default:NULL"`
	LeagueID  string `json:"leagueId"`
}

func (FifaTeam) TableName() string {
	return "generator.fifa_team"
}
