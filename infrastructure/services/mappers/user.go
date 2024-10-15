package mappers

import (
	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

func UserDtoToUserDb(userDTO dto.UserDTO) entities.User {
	return entities.User{
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		IsBlocked: userDTO.IsBlocked,
	}
}

func UserToUserDto(user entities.User) dto.UserDTO {
	return dto.UserDTO{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		IsBlocked: user.IsBlocked,
	}
}
