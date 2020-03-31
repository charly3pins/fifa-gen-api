package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func FifaLeague() fifaLeague {
	return fifaLeague{}
}

type fifaLeague struct{}

func (fifaLeague) Find(db *gorm.DB) ([]model.FifaLeague, error) {
	var res []model.FifaLeague
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
