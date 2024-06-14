package app_interfaces

import (
	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
)

type AuthUsecase interface {
	Signup(params requests.SignupReq) (entities.User, error)
	Signin(params requests.SigninReq) (string, error)
	ForgotPasswordToken(params requests.ForgotPasswordReq) error
	ResetPasswordConfirm(params requests.ResetPasswordReq) error
}
