package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func FifaPlayer() fifaPlayer {
	return fifaPlayer{}
}

type fifaPlayer struct{}

func (fifaPlayer) Find(findBy model.FifaPlayer, db *gorm.DB) ([]model.FifaPlayer, error) {
	var res []model.FifaPlayer
	if err := db.Where(findBy).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (fifaPlayer) Get(getBy model.FifaPlayer, db *gorm.DB) (model.FifaPlayer, error) {
	var f model.FifaPlayer
	if result := db.Where(getBy).First(&f); result.Error != nil {
		if result.RecordNotFound() {
			return f, nil
		}

		return f, result.Error
	}

	return f, nil
}
