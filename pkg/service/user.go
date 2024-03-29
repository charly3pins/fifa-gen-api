package service

import (
	"fmt"
	"log"

	. "github.com/charly3pins/fifa-gen-api/internal"
	"github.com/charly3pins/fifa-gen-api/pkg/model"
	repo "github.com/charly3pins/fifa-gen-api/pkg/repository"

	"github.com/jinzhu/gorm"
)

func NewUser() User {
	db, err := NewDB()
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return User{
		db: db,
	}
}

type User struct {
	db *gorm.DB
}

func (u User) Create(usr model.User) (model.User, error) {
	getBy := model.User{
		Username: usr.Username,
	}
	usrDB, err := repo.User().Get(getBy, u.db)
	if err != nil {
		log.Printf("error getting the User for Username %s:\n%s\n", usr.Username, err)
		return usr, err
	}
	if usrDB.ID != "" {
		// TODO return specific code
		return usr, fmt.Errorf("error duplicate User for Username %s", usr.Username)
	}

	usrDB, err = repo.User().Create(usr, u.db)
	if err != nil {
		log.Printf("error creating the User %+v:\n%s\n", usr, err)
		return usr, err
	}

	return usrDB, nil
}

func (u User) Get(getBy model.User) (model.User, error) {
	usr, err := repo.User().Get(getBy, u.db)
	if err != nil {
		log.Printf("error getting the User %+v:\n%s\n", getBy, err)
		return usr, err
	}
	// TODO check if user exists if not return specific code

	return usr, nil
}

func (u User) Update(usr model.User) error {
	getBy := model.User{
		ID:       usr.ID,
		Username: usr.Username,
	}
	usrDB, err := repo.User().Get(getBy, u.db)
	if err != nil {
		log.Printf("error getting the User for %+v:\n%s\n", getBy, err)
		return err
	}
	if usrDB.ID == "" {
		// TODO return specific code
		return fmt.Errorf("User for %+v not found", getBy)
	}

	if err := repo.User().Update(usr, u.db); err != nil {
		log.Printf("error updating the User %+v:\n%s\n", usr, err)
		return err
	}

	return nil
}

func (u User) Find(findBy model.User) ([]model.User, error) {
	f, err := repo.User().Find(findBy, u.db)
	if err != nil {
		log.Printf("error finding the User for Username %s:\n%s\n", findBy.Username, err)
		return f, err
	}

	return f, nil
}
