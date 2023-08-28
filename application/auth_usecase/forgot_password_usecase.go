package auth_usecase

import (
	"fmt"
	"math/rand"

	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/mailers"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
)

func (u *authUsecase) ForgotPasswordToken(params requests.ForgotPasswordReq) error {
	//Find user on DB
	user, err := u.repository.GetUserByEmail(params.Email)
	if err != nil {
		return err
	}

	//Generating Token
	token := fmt.Sprintf("%d", (100000 + rand.Intn(899999)))

	//Persisting the token
	userWithToken := entities.UserDB{
		TokenResetPassword: token,
	}
	err = u.repository.UpdateUser(user.ID, userWithToken)
	if err != nil {
		return u.configs.Exceptions.ErrForgotPassword
	}

	//Sending token by email
	ms := mailers.NewMailSender(u.configs)
	go ms.TokenForgotPassword(user.Email, token)

	return nil
}
