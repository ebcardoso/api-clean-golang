package auth_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/infrastructure/services/jwt_parser"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
)

func (u *authUsecase) Signin(params requests.SigninReq) (string, error) {
	//Find user on BD
	user, err := u.repository.GetUserByEmail(params.Email)
	if err != nil {
		return "", err
	}

	//Check if the user is blocked
	if user.IsBlocked {
		return "", u.configs.Exceptions.ErrUserIsBlocked
	}

	//Call check password
	checkPassword := entities.CheckPassword(user, params.Password)
	if !checkPassword {
		return "", u.configs.Exceptions.ErrWrongPassword
	}

	//Generate token jwt
	token, err := jwt_parser.EncodeJWT(user.ID.Hex(), u.configs.Env.JWT_KEY)
	if err != nil {
		return "", err
	}

	//Success
	return token, nil
}
