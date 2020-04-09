package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func Friend() friend {
	return friend{}
}

type friend struct{}

func (friend) Create(f model.Friend, db *gorm.DB) (model.Friend, error) {
	if db.NewRecord(f) {
		if err := db.Create(&f).Error; err != nil {
			return f, err
		}
	}

	return f, nil
}

func (friend) Get(getBy model.Friend, db *gorm.DB) (model.Friend, error) {
	var f model.Friend
	if result := db.Where(getBy).First(&f); result.Error != nil {
		if result.RecordNotFound() {
			return f, nil
		}

		return f, result.Error
	}

	return f, nil
}

func (friend) Update(f model.Friend, db *gorm.DB) error {
	if err := db.Save(&f).Error; err != nil {
		return err
	}

	return nil
}
