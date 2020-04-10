package model

type FifaPlayer struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Number     int    `json:"number"`
	PictureSrc string `json:"pictureSrc" gorm:"default:NULL"`
	TeamID     string `json:"teamId"`
}

func (FifaPlayer) TableName() string {
	return "generator.fifa_player"
}
