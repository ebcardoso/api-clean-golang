package auth_usecase

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
	"golang.org/x/crypto/bcrypt"
)

func (u *authUsecase) ResetPasswordConfirm(params requests.ResetPasswordReq) error {
	// Check if the password are the same
	if params.Password != params.PasswordConfirmation {
		return u.configs.Exceptions.ErrDifferentPassword
	}

	// Find user on DB
	user, err := u.repository.GetUserByEmail(params.Email)
	if err != nil {
		return err
	}

	// Checking token
	if user.TokenResetPassword != params.Token {
		return u.configs.Exceptions.ErrInvalidToken
	}

	// Persisting new password
	// Hashed Password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 8)
	if err != nil {
		return u.configs.Exceptions.ErrResetPassword
	}

	// User Object
	userPasswordReset := entities.User{
		TokenResetPassword: "@-@-@-@",
		Password:           string(encryptedPassword[:]),
	}

	// Save User
	err = u.repository.UpdateUser(user.ID.Hex(), userPasswordReset)
	if err != nil {
		return u.configs.Exceptions.ErrResetPassword
	}

	return nil
}
