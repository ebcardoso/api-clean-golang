package app_interfaces

import "github.com/ebcardoso/api-clean-golang/domain/entities"

type UsersUsecase interface {
	GetList() ([]entities.User, error)
	Create(entities.UserDB) (entities.UserDB, error)
	GetByID(id string) (entities.UserDB, error)
	Update(id string, user entities.UserDB) error
	Destroy(id string) error
	BlockUnblock(id string, isBlocked bool) error
}
