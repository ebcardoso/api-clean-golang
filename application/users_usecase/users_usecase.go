package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/interfaces"
	"github.com/ebcardoso/api-clean-golang/domain/repositories"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
)

type usersUsecase struct {
	repository interfaces.UsersRepository
	configs    *config.Config
}

func NewUsersUsecase(configs *config.Config) interfaces.UsersUsecase {
	return &usersUsecase{
		repository: repositories.NewUsersMongoRepository(configs),
		configs:    configs,
	}
}
