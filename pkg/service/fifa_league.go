package service

import (
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewFifaLeague() FifaLeague {
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return FifaLeague{
		db: db,
	}
}

type FifaLeague struct {
	db *gorm.DB
}

func (fl FifaLeague) Find() ([]model.FifaLeague, error) {
	f, err := repo.FifaLeague().Find(fl.db)
	if err != nil {
		log.Printf("error finding the Fifa Leagues:\n%s\n", err)
		return f, err
	}

	return f, nil
}
