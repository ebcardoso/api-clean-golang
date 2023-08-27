package exceptions

import (
	"errors"

	"github.com/ebcardoso/api-clean-golang/infrastructure/translations"
)

type Exceptions struct {
	//Auth
	ErrWrongPassword     error
	ErrDifferentPassword error
	ErrUserAlreadyExists error
	ErrSaveUser          error

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
		//Auth
		ErrWrongPassword:     errors.New(translations.Auth.Signin.Invalid),
		ErrDifferentPassword: errors.New(translations.Auth.Signup.Errors.PasswordDifferent),
		ErrUserAlreadyExists: errors.New(translations.Auth.Signup.Errors.AlreadyExists),
		ErrSaveUser:          errors.New(translations.Auth.Signup.Errors.SaveUser),

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
