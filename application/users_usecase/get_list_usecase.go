package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/mappers"
)

func (u *usersUsecase) GetList() ([]dto.UserDTO, error) {
	users, err := u.repository.ListUsers()
	if err != nil {
		return nil, err
	}

	items := []dto.UserDTO{}
	for _, user := range users {
		items = append(items, mappers.UserToUserDto(user))
	}

	return items, err
}
