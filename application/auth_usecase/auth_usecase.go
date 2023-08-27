package auth_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/interfaces"
	"github.com/ebcardoso/api-clean-golang/domain/repositories"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
)

type authUsecase struct {
	repository interfaces.UsersRepository
	configs    *config.Config
}

func NewAuthUseCase(configs *config.Config) interfaces.AuthUsecase {
	return &authUsecase{
		repository: repositories.NewUsersMongoRepository(configs),
		configs:    configs,
	}
}
