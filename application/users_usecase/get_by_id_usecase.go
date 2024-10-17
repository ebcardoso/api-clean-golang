package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/mappers"
)

func (u *usersUsecase) GetByID(id string) (dto.UserDTO, error) {
	result, err := u.repository.GetUserByID(id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return mappers.UserToUserDto(result), nil
}
