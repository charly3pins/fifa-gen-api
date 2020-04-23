package repository

import (
	"log"

	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func UserGroup() userGroup {
	return userGroup{}
}

type userGroup struct{}

func (userGroup) Create(ug model.UserGroup, db *gorm.DB) (model.UserGroup, error) {
	if db.NewRecord(ug) {
		if err := db.Create(&ug).Error; err != nil {
			return ug, err
		}
	}

	return ug, nil
}

func (userGroup) Delete(ug model.UserGroup, db *gorm.DB) error {
	log.Printf("delete ug %+v", ug)
	return db.Where("user_id = ? AND group_id = ?", ug.UserID, ug.GroupID).Delete(model.UserGroup{}).Error
}

func (userGroup) Get(getBy model.UserGroup, db *gorm.DB) (model.UserGroup, error) {
	var ug model.UserGroup
	if result := db.Where(getBy).First(&ug); result.Error != nil {
		if result.RecordNotFound() {
			return ug, nil
		}

		return ug, result.Error
	}

	return ug, nil
}

func (userGroup) FindMembers(groupID string, db *gorm.DB) ([]model.Member, error) {
	var res []model.Member

	if err := db.Raw("SELECT u.id, u.name, u.username, u.active, u.profile_picture, ug.is_admin FROM "+model.User{}.TableName()+" u "+
		"JOIN "+model.UserGroup{}.TableName()+" ug ON u.id = ug.user_id "+
		"WHERE ug.group_id = ?", groupID).
		Find(&res).
		Error; err != nil {
	}

	return res, nil
}

func (userGroup) Update(ug model.UserGroup, db *gorm.DB) error {
	if err := db.Model(&ug).Update("is_admin", ug.IsAdmin).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
