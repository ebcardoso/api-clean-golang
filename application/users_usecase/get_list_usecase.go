package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

func (u *usersUsecase) GetList() ([]entities.User, error) {
	users, err := u.repository.ListUsers()
	if err != nil {
		return nil, err
	}

	items := []entities.User{}
	for _, user := range users {
		items = append(items, user.MapUserDB())
	}

	return items, err
}
