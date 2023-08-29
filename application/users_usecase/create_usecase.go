package users_usecase

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

func (u *usersUsecase) Create(user entities.UserDB) (entities.UserDB, error) {
	//Check if the user already exists
	_, err := u.repository.GetUserByEmail(user.Email)
	if err != nil {
		if errors.Is(err, u.configs.Exceptions.ErrUserGet) {
			return entities.UserDB{}, u.configs.Exceptions.ErrSaveUser
		}
	} else {
		return entities.UserDB{}, u.configs.Exceptions.ErrUserAlreadyExists
	}

	//Hashed Password
	password := fmt.Sprintf("%d", (10000000 + rand.Intn(89999999)))
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return entities.UserDB{}, u.configs.Exceptions.ErrSaveUser
	}

	user.Password = string(encryptedPassword[:])
	user.IsBlocked = false

	//Persisting User
	result, err := u.repository.CreateUser(user)
	if err != nil {
		return entities.UserDB{}, u.configs.Exceptions.ErrSaveUser
	}

	return result, nil
}
