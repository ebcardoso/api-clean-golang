package auth_usecase

import (
	"github.com/ebcardoso/api-clean-golang/application/app_interfaces"
	"github.com/ebcardoso/api-clean-golang/domain/repository"
	"github.com/ebcardoso/api-clean-golang/domain/repository_interfaces"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
)

type authUsecase struct {
	repository repository_interfaces.UsersRepository
	configs    *config.Config
}

func NewAuthUseCase(configs *config.Config) app_interfaces.AuthUsecase {
	return &authUsecase{
		repository: repository.NewUsersMongoRepository(configs),
		configs:    configs,
	}
}
