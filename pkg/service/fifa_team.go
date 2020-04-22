package service

import (
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewFifaTeam() FifaTeam {
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return FifaTeam{
		db: db,
	}
}

type FifaTeam struct {
	db *gorm.DB
}

func (ft FifaTeam) Get(getBy model.FifaTeam) (model.FifaTeam, error) {
	f, err := repo.FifaTeam().Get(getBy, ft.db)
	if err != nil {
		log.Printf("error getting the Fifa Team %+v:\n%s\n", getBy, err)
		return f, err
	}

	return f, nil
}

func (ft FifaTeam) Find(findBy model.FifaTeam) ([]model.FifaTeam, error) {
	f, err := repo.FifaTeam().Find(findBy, ft.db)
	if err != nil {
		log.Printf("error finding the Fifa Teams:\n%s\n", err)
		return f, err
	}

	return f, nil
}
