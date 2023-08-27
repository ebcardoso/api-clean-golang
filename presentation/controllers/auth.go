package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ebcardoso/api-clean-golang/application/auth_usecase"
	"github.com/ebcardoso/api-clean-golang/domain/interfaces"
	"github.com/ebcardoso/api-clean-golang/domain/repositories"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
	"github.com/ebcardoso/api-clean-golang/presentation/response"
)

type Auth struct {
	repository interfaces.UsersRepository
	usecase    interfaces.AuthUsecase
	configs    *config.Config
}

func NewAuth(configs *config.Config) *Auth {
	return &Auth{
		repository: repositories.NewUsersMongoRepository(configs),
		usecase:    auth_usecase.NewAuthUseCase(configs),
		configs:    configs,
	}
}

func (c *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	var params requests.SignupReq
	json.NewDecoder(r.Body).Decode(&params)

	result, err := c.usecase.Signup(params)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrDifferentPassword) {
			status = http.StatusBadRequest
		} else if errors.Is(err, c.configs.Exceptions.ErrUserAlreadyExists) {
			status = http.StatusUnprocessableEntity
		} else {
			status = http.StatusInternalServerError
		}

		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	response.JsonRes(w, result, http.StatusCreated)
}

func (c *Auth) Signin(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	var params requests.SigninReq
	json.NewDecoder(r.Body).Decode(&params)

	token, err := c.usecase.Signin(params)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, c.configs.Exceptions.ErrUserIsBlocked) {
			status = http.StatusUnauthorized
		} else if errors.Is(err, c.configs.Exceptions.ErrWrongPassword) {
			status = http.StatusUnauthorized
		} else {
			status = http.StatusInternalServerError
		}

		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	output["message"] = c.configs.Translations.Auth.Signin.Success
	output["accessToken"] = token
	response.JsonRes(w, output, http.StatusOK)
}

func (c *Auth) CheckToken(w http.ResponseWriter, r *http.Request) {

}
