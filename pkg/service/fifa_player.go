package service

import (
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewFifaPlayer() FifaPlayer {
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return FifaPlayer{
		db: db,
	}
}

type FifaPlayer struct {
	db *gorm.DB
}

func (ft FifaPlayer) Get(getBy model.FifaPlayer) (model.FifaPlayer, error) {
	f, err := repo.FifaPlayer().Get(getBy, ft.db)
	if err != nil {
		log.Printf("error getting the Fifa Player %+v:\n%s\n", getBy, err)
		return f, err
	}

	return f, nil
}

func (ft FifaPlayer) Find(findBy model.FifaPlayer) ([]model.FifaPlayer, error) {
	f, err := repo.FifaPlayer().Find(findBy, ft.db)
	if err != nil {
		log.Printf("error finding the Fifa Players:\n%s\n", err)
		return f, err
	}

	return f, nil
}
