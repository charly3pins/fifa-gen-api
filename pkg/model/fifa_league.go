package model

type FifaLeague struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (FifaLeague) TableName() string {
	return "generator.fifa_league"
}
