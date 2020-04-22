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
	db, err := NewDB()
	db.LogMode(true)
	if err != nil {
		log.Fatal("error creating new DB", err)
	}
	return Group{
		db: db,
	}
}

type Group struct {
	db *gorm.DB
}

func (g Group) Create(group model.GroupComplete) (model.Group, error) {
	getBy := model.Group{
		Name: group.Group.Name,
	}
	groupDB, err := repo.Group().Get(getBy, g.db)
	if err != nil {
		log.Printf("error getting the Group for Name %s:\n%s\n", group.Group.Name, err)
		return groupDB, err
	}
	if groupDB.ID != "" {
		// TODO return specific code
		return groupDB, fmt.Errorf("error duplicate Group for Name %s", group.Group.Name)
	}

	// Create group and store all members or rollback
	tx := g.db.Begin()
	defer tx.Rollback()
	groupDB, err = repo.Group().Create(group.Group, tx)
	if err != nil {
		log.Printf("error creating the Group %+v:\n%s\n", group, err)
		return groupDB, err
	}

	for _, member := range group.Members {
		userGroup := model.UserGroup{
			GroupID: groupDB.ID,
			UserID:  member.User.ID,
			IsAdmin: member.IsAdmin,
		}
		if _, err = repo.UserGroup().Create(userGroup, tx); err != nil {
			log.Printf("error creating the UserGroup %+v:\n%s\n", userGroup, err)
			return groupDB, err
		}
	}
	tx.Commit()

	return groupDB, nil
}

func (g Group) Find(findBy model.Group) ([]model.GroupComplete, error) {
	// TODO find user_group where user_id = id provided and join with groups and return them
	return nil, nil
}

func (g Group) Get(getBy model.Group) (model.GroupComplete, error) {
	var groupComplete model.GroupComplete

	group, err := repo.Group().Get(getBy, g.db)
	if err != nil {
		log.Printf("error getting the Group %+v:\n%s\n", getBy, err)
		return groupComplete, err
	}
	// TODO check if group exists if not return specific code
	groupComplete.Group = group

	// find members for that groupID
	members, err := repo.UserGroup().Find(group.ID, g.db)
	if err != nil {
		log.Printf("error finding the UserGroup members for groupID %s:\n%s\n", group.ID, err)
		return groupComplete, err
	}
	groupComplete.Members = members

	return groupComplete, nil
}

func (g Group) Update(group model.GroupComplete) error {
	getBy := model.Group{
		ID: group.Group.ID,
	}
	groupDB, err := repo.Group().Get(getBy, g.db)
	if err != nil {
		log.Printf("error getting the Group for ID %s:\n%s\n", group.Group.ID, err)
		return err
	}
	if groupDB.ID == "" {
		// TODO return specific code
		return fmt.Errorf("Group for ID %s not found", group.Group.ID)
	}

	// Update group and all members or none
	tx := g.db.Begin()
	defer tx.Rollback()
	if err := repo.Group().Update(group.Group, tx); err != nil {
		log.Printf("error updating the Group %+v:\n%s\n", group, err)
		return err
	}

	// find stored members for that groupID
	members, err := repo.UserGroup().Find(group.Group.ID, g.db)
	if err != nil {
		log.Printf("error finding the old members for groupID %s:\n%s\n", group.Group.ID, err)
		return err
	}

	// Loop over all old members and update the ones found it in the new list
	// The old members that are not present in the new list, has to be removed
	// In order to create the new ones, store the ones that are updated,
	// and loop over them, and find the ones that are not there, for create them
	var (
		foundMembers []model.Member
		userGroup    model.UserGroup
		isFound      bool
	)
	for _, oldMember := range members {
		userGroup = model.UserGroup{
			GroupID: groupDB.ID,
			UserID:  oldMember.User.ID,
			IsAdmin: oldMember.IsAdmin,
		}
		isFound = false
		for _, newMember := range group.Members {
			// If the member is on the previous list, it has to be updated
			if oldMember.ID == newMember.ID {
				if err = repo.UserGroup().Update(userGroup, tx); err != nil {
					log.Printf("error updating the UserGroup %+v:\n%s\n", userGroup, err)
					return err
				}
				foundMembers = append(foundMembers, newMember)
				isFound = true
				break
			}
			isFound = false
		}
		// If old member is not in the new member list, it has to be removed
		if !isFound {
			if err = repo.UserGroup().Delete(userGroup, tx); err != nil {
				log.Printf("error deleting the UserGroup %+v:\n%s\n", userGroup, err)
				return err
			}
		}
	}
	// Create new ones
	for _, newMember := range group.Members {
		userGroup = model.UserGroup{
			GroupID: groupDB.ID,
			UserID:  newMember.User.ID,
			IsAdmin: newMember.IsAdmin,
		}
		isFound = false
		for _, alreadyUpdatedMember := range foundMembers {
			// If the new member is in the list, it's already updated, so skip to the next one
			if alreadyUpdatedMember.ID == newMember.ID {
				isFound = true
				break
			}
			isFound = false
		}
		// If the new member is not found in the members already updated, it has to be created
		if !isFound {
			if _, err = repo.UserGroup().Create(userGroup, tx); err != nil {
				log.Printf("error creating the UserGroup %+v:\n%s\n", userGroup, err)
				return err
			}
		}
	}
	tx.Commit()

	return nil
}
