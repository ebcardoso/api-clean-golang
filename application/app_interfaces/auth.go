package app_interfaces

import (
	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
)

type AuthUsecase interface {
	Signup(params requests.SignupReq) (dto.UserDTO, error)
	Signin(params requests.SigninReq) (string, error)
	ForgotPasswordToken(params requests.ForgotPasswordReq) error
	ResetPasswordConfirm(params requests.ResetPasswordReq) error
}
