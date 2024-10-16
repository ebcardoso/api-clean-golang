package repository_interfaces

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

type UsersRepository interface {
	ListUsers() ([]entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	GetUserByID(id string) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	UpdateUser(id string, user entities.User) error
	DestroyUser(id string) error
	BlockUnblockUser(id string, isBlocked bool) error
}
