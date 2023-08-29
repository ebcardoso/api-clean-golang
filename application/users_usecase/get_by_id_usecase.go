package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
)

func (u *usersUsecase) GetByID(id string) (entities.UserDB, error) {
	result, err := u.repository.GetUserByID(id)
	if err != nil {
		return entities.UserDB{}, err
	}

	return result, nil
}
