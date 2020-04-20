package repository

import (
	"fmt"
	"strconv"

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

func (friendship) Find(userID, filter string, db *gorm.DB) ([]model.User, error) {
	var users []model.User
	var where string
	switch filter {
	case model.FilterRequested:
		// Requested friendships sent
		where = "WHERE ordered_fs.user_one_id = ? AND ordered_fs.status = " + strconv.Itoa(model.StatusCodeRequested) + " AND ordered_fs.action_user_id = ?"
		break
	case model.FilterPending:
		// Pending requets received
		where = "WHERE ordered_fs.user_one_id = ? AND ordered_fs.status = " + strconv.Itoa(model.StatusCodeRequested) + " AND ordered_fs.action_user_id <> ?"
		break
	case model.FilterFriends:
		// Friends
		where = "WHERE (ordered_fs.user_one_id = ? OR ordered_fs.user_two_id = ?) AND ordered_fs.status = " + strconv.Itoa(model.StatusCodeAccepted) + " GROUP BY (u.id)"
		break
	default:
		return users, fmt.Errorf("filter %s not implemented", filter)
	}

	if err := db.Raw("SELECT u.id, u.name, u.username, u.active, u.profile_picture FROM "+
		"("+
		"( SELECT r.user_one_id, r.user_two_id, r.status, r.action_user_id FROM "+model.Friendship{}.TableName()+" AS r where r.user_one_id = ?)"+
		"UNION"+
		"( SELECT p.user_two_id, p.user_one_id, p.status, p.action_user_id FROM "+model.Friendship{}.TableName()+" AS p where p.user_two_id = ?)"+
		") AS ordered_fs "+
		"JOIN "+model.User{}.TableName()+" u ON u.id = ordered_fs.user_two_id "+
		where, userID, userID, userID, userID).
		Find(&users).
		Error; err != nil {
	}

	return users, nil
}

func (friendship) Get(getBy model.Friendship, db *gorm.DB) (model.Friendship, error) {
	var f model.Friendship
	if result := db.
		Where("(user_one_id = ? AND user_two_id = ?) OR (user_one_id = ? AND user_two_id = ?)",
			getBy.UserOneID, getBy.UserTwoID, getBy.UserTwoID, getBy.UserOneID).
		First(&f); result.Error != nil {
		if result.RecordNotFound() {
			return f, nil
		}

		return f, result.Error
	}

	return f, nil
}

func (friendship) Update(f model.Friendship, db *gorm.DB) error {
	if err := db.Model(&f).
		Where("user_one_id = ? AND user_two_id = ?", f.UserOneID, f.UserTwoID).
		UpdateColumns(model.Friendship{Status: f.Status, ActionUserID: f.ActionUserID}).Error; err != nil {
		return err
	}

	return nil
}
