package repository

import (
	"github.com/charly3pins/fifa-gen-api/pkg/model"

	"github.com/jinzhu/gorm"
)

func FifaLeague() fifaLeague {
	return fifaLeague{}
}

type fifaLeague struct{}

// func (fl fifaLeague) Create(f model.FifaLeague, db *gorm.DB) (model.FifaLeague, error) {
// 	if db.NewRecord(f) {
// 		if err := db.Create(&f).Error; err != nil {
// 			return f, err
// 		}
// 	}

// 	return f, nil
// }

func (fl fifaLeague) Find(db *gorm.DB) ([]model.FifaLeague, error) {
	var res []model.FifaLeague
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// func (fl fifaLeague) Update(f model.FifaLeague, db *gorm.DB) error {
// 	if err := db.Save(&f).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
