package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func FifaTeam() fifaTeam {
	return fifaTeam{}
}

type fifaTeam struct{}

func (fifaTeam) Find(findBy model.FifaTeam, db *gorm.DB) ([]model.FifaTeam, error) {
	var res []model.FifaTeam
	if err := db.Where(findBy).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (fifaTeam) Get(getBy model.FifaTeam, db *gorm.DB) (model.FifaTeam, error) {
	var f model.FifaTeam
	if result := db.Where(getBy).First(&f); result.Error != nil {
		if result.RecordNotFound() {
			return f, nil
		}

		return f, result.Error
	}

	return f, nil
}
