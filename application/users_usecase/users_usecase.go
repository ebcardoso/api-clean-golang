package users_usecase

import (
	"github.com/ebcardoso/api-clean-golang/application/app_interfaces"
	"github.com/ebcardoso/api-clean-golang/domain/repository"
	"github.com/ebcardoso/api-clean-golang/domain/repository_interfaces"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
)

type usersUsecase struct {
	repository repository_interfaces.UsersRepository
	configs    *config.Config
}

func NewUsersUsecase(configs *config.Config) app_interfaces.UsersUsecase {
	return &usersUsecase{
		repository: repository.NewUsersMongoRepository(configs),
		configs:    configs,
	}
}
