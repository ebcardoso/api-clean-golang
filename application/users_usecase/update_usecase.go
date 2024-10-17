package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/mappers"
)

func (u *usersUsecase) Update(id string, userDTO dto.UserDTO) error {
	user := mappers.UserDtoToUserDb(userDTO)

	err := u.repository.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}
