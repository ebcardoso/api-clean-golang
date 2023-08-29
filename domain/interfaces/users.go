package interfaces

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

type UsersRepository interface {
	ListUsers() ([]entities.User, error)
	CreateUser(input entities.UserDB) (entities.UserDB, error)
	GetUserByID(id string) (entities.UserDB, error)
	GetUserByEmail(email string) (entities.UserDB, error)
	UpdateUser(id string, input entities.UserDB) error
	DestroyUser(id string) error
	BlockUnblockUser(id string, isBlocked bool) error
}

type UsersUsecase interface {
	GetList() ([]entities.User, error)
	Create(entities.UserDB) (entities.UserDB, error)
	GetByID(id string) (entities.UserDB, error)
	Update(id string, user entities.UserDB) error
	Destroy(id string) error
	BlockUnblock(id string, isBlocked bool) error
}
