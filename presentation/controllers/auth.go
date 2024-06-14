package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ebcardoso/api-clean-golang/application/app_interfaces"
	"github.com/ebcardoso/api-clean-golang/application/auth_usecase"
	"github.com/ebcardoso/api-clean-golang/domain/repository"
	"github.com/ebcardoso/api-clean-golang/domain/repository_interfaces"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
	"github.com/ebcardoso/api-clean-golang/presentation/requests_contract"
	"github.com/ebcardoso/api-clean-golang/presentation/response"
)

type Auth struct {
	repository repository_interfaces.UsersRepository
	usecase    app_interfaces.AuthUsecase
	configs    *config.Config
}

func NewAuth(configs *config.Config) *Auth {
	return &Auth{
		repository: repository.NewUsersMongoRepository(configs),
		usecase:    auth_usecase.NewAuthUseCase(configs),
		configs:    configs,
	}
}

func (c *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	//Post Params
	var params requests.SignupReq
	json.NewDecoder(r.Body).Decode(&params)

	//Validating Inputs
	contract, err := requests_contract.Validate(params)
	if err != nil {
		output["message"] = c.configs.Translations.Errors.InvalidParams
		output["errors"] = contract
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

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

	//Post Params
	var params requests.SigninReq
	json.NewDecoder(r.Body).Decode(&params)

	//Validating Inputs
	contract, err := requests_contract.Validate(params)
	if err != nil {
		output["message"] = c.configs.Translations.Errors.InvalidParams
		output["errors"] = contract
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

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
	output["access_token"] = token
	response.JsonRes(w, output, http.StatusOK)
}

func (c *Auth) CheckToken(w http.ResponseWriter, r *http.Request) {

}

func (c *Auth) ForgotPasswordToken(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	//Post Params
	var params requests.ForgotPasswordReq
	json.NewDecoder(r.Body).Decode(&params)

	//Validating Inputs
	contract, err := requests_contract.Validate(params)
	if err != nil {
		output["message"] = c.configs.Translations.Errors.InvalidParams
		output["errors"] = contract
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	err = c.usecase.ForgotPasswordToken(params)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	//Success Response
	output["message"] = c.configs.Translations.Auth.ForgotPasswordToken.Success
	response.JsonRes(w, output, http.StatusOK)
}

func (c *Auth) ResetPasswordConfirm(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	//Post Params
	var params requests.ResetPasswordReq
	json.NewDecoder(r.Body).Decode(&params)

	//Validating Inputs
	contract, err := requests_contract.Validate(params)
	if err != nil {
		output["message"] = c.configs.Translations.Errors.InvalidParams
		output["errors"] = contract
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	err = c.usecase.ResetPasswordConfirm(params)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, c.configs.Exceptions.ErrDifferentPassword) {
			status = http.StatusUnprocessableEntity
		} else if errors.Is(err, c.configs.Exceptions.ErrInvalidToken) {
			status = http.StatusForbidden
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	//Success Response
	output["message"] = c.configs.Translations.Auth.ResetPasswordConfirm.Success
	response.JsonRes(w, output, http.StatusOK)
}
