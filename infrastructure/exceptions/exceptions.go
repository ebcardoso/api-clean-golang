package exceptions

import (
	"errors"

	"github.com/ebcardoso/api-clean-golang/infrastructure/translations"
)

type Exceptions struct {
	//General
	ErrParseId       error
	ErrInvalidParams error

	//Auth
	ErrWrongPassword     error
	ErrDifferentPassword error
	ErrUserAlreadyExists error
	ErrSaveUser          error
	ErrForgotPassword    error
	ErrResetPassword     error
	ErrInvalidToken      error

	//Users
	ErrUserNotFound  error
	ErrUserList      error
	ErrUserFetch     error
	ErrUserGet       error
	ErrUserCreate    error
	ErrUserUpdate    error
	ErrUserDestroy   error
	ErrUserBlock     error
	ErrUserUnblock   error
	ErrUserIsBlocked error
}

func NewExceptions(translations *translations.Translations) *Exceptions {
	return &Exceptions{
		//General
		ErrParseId:       errors.New(translations.Errors.ParseId),
		ErrInvalidParams: errors.New(translations.Errors.InvalidParams),

		//Auth
		ErrWrongPassword:     errors.New(translations.Auth.Signin.Invalid),
		ErrDifferentPassword: errors.New(translations.Auth.Signup.Errors.PasswordDifferent),
		ErrUserAlreadyExists: errors.New(translations.Auth.Signup.Errors.AlreadyExists),
		ErrSaveUser:          errors.New(translations.Auth.Signup.Errors.SaveUser),
		ErrForgotPassword:    errors.New(translations.Auth.ForgotPasswordToken.Errors.Default),
		ErrResetPassword:     errors.New(translations.Auth.ResetPasswordConfirm.Errors.Default),
		ErrInvalidToken:      errors.New(translations.Auth.ResetPasswordConfirm.Errors.InvalidToken),

		//User
		ErrUserNotFound:  errors.New(translations.Users.Errors.NotFound),
		ErrUserList:      errors.New(translations.Users.List.Errors),
		ErrUserFetch:     errors.New(translations.Users.Fetch.Errors),
		ErrUserGet:       errors.New(translations.Users.Load.Errors),
		ErrUserCreate:    errors.New(translations.Users.Create.Errors),
		ErrUserUpdate:    errors.New(translations.Users.Update.Errors),
		ErrUserDestroy:   errors.New(translations.Users.Destroy.Errors),
		ErrUserBlock:     errors.New(translations.Users.Block.Errors),
		ErrUserUnblock:   errors.New(translations.Users.Unblock.Errors),
		ErrUserIsBlocked: errors.New(translations.Auth.Signin.UserBlocked),
	}
}
