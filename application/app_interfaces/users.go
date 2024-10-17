package app_interfaces

import "github.com/ebcardoso/api-clean-golang/application/dto"

type UsersUsecase interface {
	GetList() ([]dto.UserDTO, error)
	Create(userDTO dto.UserDTO) (dto.UserDTO, error)
	GetByID(id string) (dto.UserDTO, error)
	Update(id string, user dto.UserDTO) error
	Destroy(id string) error
	BlockUnblock(id string, isBlocked bool) error
}
