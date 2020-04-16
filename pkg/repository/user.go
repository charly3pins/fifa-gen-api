package repository

import (
	"strings"

	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func User() user {
	return user{}
}

type user struct{}

func (user) Create(u model.User, db *gorm.DB) (model.User, error) {
	if db.NewRecord(u) {
		if err := db.Create(&u).Error; err != nil {
			return u, err
		}
	}

	return u, nil
}

func (user) Get(getBy model.User, db *gorm.DB) (model.User, error) {
	var u model.User
	if result := db.Where(getBy).First(&u); result.Error != nil {
		if result.RecordNotFound() {
			return u, nil
		}

		return u, result.Error
	}

	return u, nil
}

func (user) Update(u model.User, db *gorm.DB) error {
	if err := db.Model(&u).
		Where("id = ?", u.ID).
		UpdateColumns(model.User{Name: u.Name}).Error; err != nil {
		return err
	}

	return nil
}

func (user) Find(findBy model.User, db *gorm.DB) ([]model.User, error) {
	var res []model.User
	// TODO improve the sanytize for the %username%
	if err := db.Where("UPPER(username) LIKE ?", "%"+strings.ToUpper(findBy.Username)+"%").
		Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
