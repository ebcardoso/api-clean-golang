package interfaces

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

type UsersRepository interface {
	CreateUser(input entities.UserDB) (entities.User, error)
	GetUserByEmail(email string) (entities.UserDB, error)
}
