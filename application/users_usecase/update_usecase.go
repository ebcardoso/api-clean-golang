package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

func (u *usersUsecase) Update(id string, user entities.UserDB) error {
	err := u.repository.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}
