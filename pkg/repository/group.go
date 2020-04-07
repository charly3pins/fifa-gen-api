package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func Group() group {
	return group{}
}

type group struct{}

func (group) Find(findBy model.Group, db *gorm.DB) ([]model.Group, error) {
	var res []model.Group
	if err := db.Where(findBy).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (group) Get(getBy model.Group, db *gorm.DB) (model.Group, error) {
	var g model.Group
	if result := db.Where(getBy).First(&g); result.Error != nil {
		if result.RecordNotFound() {
			return g, nil
		}

		return g, result.Error
	}

	return g, nil
}

func (group) Create(gr model.Group, db *gorm.DB) (model.Group, error) {
	if db.NewRecord(gr) {
		if err := db.Create(&gr).Error; err != nil {
			return gr, err
		}
	}

	return gr, nil
}

func (group) Update(gr model.Group, db *gorm.DB) error {
	if err := db.Save(&gr).Error; err != nil {
		return err
	}

	return nil
}

