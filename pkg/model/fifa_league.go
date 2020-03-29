package model

import "github.com/jinzhu/gorm"

type FifaLeague struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (fl FifaLeague) Create(f *FifaLeague, db *gorm.DB) error {
	if db.NewRecord(f) {
		if err := db.Create(f).Error; err != nil {
			return err
		}
	}

	return nil
}

func (fl FifaLeague) Find(db *gorm.DB) ([]*FifaLeague, error) {
	var res []*FifaLeague
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (fl FifaLeague) Get(getBy *FifaLeague, db *gorm.DB) (*FifaLeague, error) {
	f := new(FifaLeague)
	if result := db.Where(getBy).First(f); result.Error != nil {
		if result.RecordNotFound() {
			return nil, nil
		}

		return nil, result.Error
	}

	return f, nil
}

func (fl FifaLeague) Update(f *FifaLeague, db *gorm.DB) error {
	if err := db.Save(f).Error; err != nil {
		return err
	}

	return nil
}
