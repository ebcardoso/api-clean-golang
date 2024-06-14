package repository_interfaces

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
