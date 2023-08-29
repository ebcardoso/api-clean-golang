package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

func (u *usersUsecase) GetList() ([]entities.User, error) {
	items, err := u.repository.ListUsers()
	if err != nil {
		return nil, err
	}
	return items, err
}
