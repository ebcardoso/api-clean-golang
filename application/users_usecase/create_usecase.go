package users_usecase

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/mappers"
	"golang.org/x/crypto/bcrypt"
)

func (u *usersUsecase) Create(userDTO dto.UserDTO) (dto.UserDTO, error) {
	//Check if the user already exists
	_, err := u.repository.GetUserByEmail(userDTO.Email)
	if err != nil {
		if errors.Is(err, u.configs.Exceptions.ErrUserGet) {
			return dto.UserDTO{}, u.configs.Exceptions.ErrSaveUser
		}
	} else {
		return dto.UserDTO{}, u.configs.Exceptions.ErrUserAlreadyExists
	}

	//Hashed Password
	password := fmt.Sprintf("%d", (10000000 + rand.Intn(89999999)))
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return dto.UserDTO{}, u.configs.Exceptions.ErrSaveUser
	}

	user := mappers.UserDtoToUserDb(userDTO)
	user.Password = string(encryptedPassword[:])
	user.IsBlocked = true

	//Persisting User
	result, err := u.repository.CreateUser(user)
	if err != nil {
		return dto.UserDTO{}, u.configs.Exceptions.ErrSaveUser
	}

	return mappers.UserToUserDto(result), nil
}
