package auth_usecase

import (
	"errors"

	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/mappers"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
	"golang.org/x/crypto/bcrypt"
)

func (u *authUsecase) Signup(params requests.SignupReq) (dto.UserDTO, error) {
	//Check if the password and confirmation are the same
	if params.Password != params.PasswordConfirmation {
		return dto.UserDTO{}, u.configs.Exceptions.ErrDifferentPassword
	}

	//Check if the user already exists
	_, err := u.repository.GetUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, u.configs.Exceptions.ErrUserGet) {
			return dto.UserDTO{}, u.configs.Exceptions.ErrSaveUser
		}
	} else {
		return dto.UserDTO{}, u.configs.Exceptions.ErrUserAlreadyExists
	}

	//Hashed Password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 8)
	if err != nil {
		return dto.UserDTO{}, u.configs.Exceptions.ErrSaveUser
	}

	//Save User
	user := entities.User{
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(encryptedPassword[:]),
		IsBlocked: false,
	}
	result, err := u.repository.CreateUser(user)
	if err != nil {
		return dto.UserDTO{}, u.configs.Exceptions.ErrSaveUser
	}

	return mappers.UserToUserDto(result), nil
}
