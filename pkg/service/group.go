package service

import (
	"fmt"
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewGroup() Group {
	// Database
	db, err := NewDB()
	if err != nil {
		log.Fatal("Error creating new DB", err)
	}
	return Group{
		db: db,
	}
}

type Group struct {
	db *gorm.DB
}

func (g Group) Get(getBy model.Group) (model.Group, error) {
	f, err := repo.Group().Get(getBy, g.db)
	if err != nil {
		log.Printf("error getting the Group %+v:\n%s\n", getBy, err)
		return f, err
	}

	return f, nil
}

func (g Group) Find(findBy model.Group) ([]model.Group, error) {
	f, err := repo.Group().Find(findBy, g.db)
	if err != nil {
		log.Printf("error finding the Group:\n%s\n", err)
		return f, err
	}

	return f, nil
}

func (g Group) Create(gro model.Group) (model.Group, error) {
	getBy := model.Group{
		ID: gro.ID,
	}
	grDB, err := repo.Group().Get(getBy, g.db)
	if err != nil {
		log.Printf("error getting the Group for ID %s:\n%s\n", gro.ID, err)
		return gro, err
	}
	if grDB.ID != "" {
		// TODO return specific code
		return gro, fmt.Errorf("error duplicate Group for ID %s", gro.ID)
	}

	grDB, err = repo.Group().Create(gro, g.db)
	if err != nil {
		log.Printf("error creating the Group %+v::\n%s\n", gro, err)
		return gro, err
	}

	return grDB, nil
}

func (g Group) Update(gro model.Group) error {
	getBy := model.Group{
		ID: gro.ID,
	}
	grDB, err := repo.Group().Get(getBy, g.db)
	if err != nil {
		log.Printf("error getting the Group for ID %s:\n%s\n", gro.ID, err)
		return err
	}
	if grDB.ID == "" {
		return fmt.Errorf("Group for ID %s not found", gro.ID)
	}

	if err := repo.Group().Update(gro, g.db); err != nil {
		log.Printf("error updating the Group %+v:\n%s\n", gro, err)
		return err
	}

	return nil
}
