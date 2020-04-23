package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func Group() group {
	return group{}
}

type group struct{}

func (group) Create(g model.Group, db *gorm.DB) (model.Group, error) {
	if db.NewRecord(g) {
		if err := db.Create(&g).Error; err != nil {
			return g, err
		}
	}

	return g, nil
}

func (group) Find(userID string, db *gorm.DB) ([]model.Group, error) {
	var res []model.Group

	if err := db.Raw("SELECT g.id, g.name FROM "+model.Group{}.TableName()+" g "+
		"JOIN "+model.UserGroup{}.TableName()+" ug ON g.id = ug.group_id "+
		"WHERE ug.user_id = ?", userID).
		Find(&res).
		Error; err != nil {
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

func (group) Update(g model.Group, db *gorm.DB) error {
	if err := db.Model(&g).
		UpdateColumns(model.Group{Name: g.Name}).Error; err != nil {
		return err
	}

	return nil
}
