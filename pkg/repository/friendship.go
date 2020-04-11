package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func Friendship() friendship {
	return friendship{}
}

type friendship struct{}

func (friendship) Create(f model.Friendship, db *gorm.DB) (model.Friendship, error) {
	if db.NewRecord(f) {
		if err := db.Create(&f).Error; err != nil {
			return f, err
		}
	}

	return f, nil
}

func (friendship) Get(getBy model.Friendship, db *gorm.DB) (model.Friendship, error) {
	var f model.Friendship
	if result := db.Where("(user_one_id = ? AND user_two_id = ?) OR (user_one_id = ? AND user_two_id = ?)", getBy.UserOneID, getBy.UserTwoID, getBy.UserTwoID, getBy.UserOneID).
		First(&f); result.Error != nil {
		if result.RecordNotFound() {
			return f, nil
		}

		return f, result.Error
	}

	return f, nil
}

func (friendship) FriendsList(userID string, db *gorm.DB) ([]model.Friendship, error) {
	var res []model.Friendship
	if err := db.Where("(user_one_id = ? OR user_two_id = ?) AND status = ?", userID, userID, model.StatusCodeAccepted).
		Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (friendship) PendingRequests(userID string, db *gorm.DB) ([]model.Friendship, error) {
	var res []model.Friendship
	if err := db.Where("(user_one_id = ? OR user_two_id = ?) AND status = ? AND action_user_id <> ?", userID, userID, model.StatusCodePending, userID).
		Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (friendship) Update(f model.Friendship, db *gorm.DB) error {
	if err := db.Save(&f).Error; err != nil {
		return err
	}

	return nil
}
