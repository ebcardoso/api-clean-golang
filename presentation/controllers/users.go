package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ebcardoso/api-clean-golang/application/app_interfaces"
	"github.com/ebcardoso/api-clean-golang/application/dto"
	"github.com/ebcardoso/api-clean-golang/application/users_usecase"
	"github.com/ebcardoso/api-clean-golang/domain/repository"
	"github.com/ebcardoso/api-clean-golang/domain/repository_interfaces"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
	"github.com/ebcardoso/api-clean-golang/presentation/requests"
	"github.com/ebcardoso/api-clean-golang/presentation/response"
	"github.com/go-chi/chi/v5"
)

type Users struct {
	repository repository_interfaces.UsersRepository
	usecase    app_interfaces.UsersUsecase
	configs    *config.Config
}

func NewUsers(configs *config.Config) *Users {
	return &Users{
		repository: repository.NewUsersMongoRepository(configs),
		usecase:    users_usecase.NewUsersUsecase(configs),
		configs:    configs,
	}
}

func (c *Users) GetList(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	users, err := c.usecase.GetList()
	if err != nil {
		output["message"] = err.Error()
		response.JsonRes(w, output, http.StatusInternalServerError)
		return
	}

	output["content"] = users
	response.JsonRes(w, output, http.StatusOK)
}

func (c *Users) Create(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	var params dto.UserDTO
	json.NewDecoder(r.Body).Decode(&params)

	user, err := c.usecase.Create(params)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserAlreadyExists) {
			status = http.StatusUnprocessableEntity
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	response.JsonRes(w, user, http.StatusOK)
}

func (c *Users) GetByID(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")

	user, err := c.usecase.GetByID(id)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrParseId) {
			status = http.StatusBadRequest
		} else if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	response.JsonRes(w, user, http.StatusOK)
}

func (c *Users) Update(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")

	var params requests.UsersUpdateReq
	json.NewDecoder(r.Body).Decode(&params)

	user := dto.UserDTO{
		Name: params.Name,
	}

	err := c.usecase.Update(id, user)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, c.configs.Exceptions.ErrInvalidParams) {
			status = http.StatusBadRequest
		} else if errors.Is(err, c.configs.Exceptions.ErrParseId) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	result := user
	result.ID = id
	response.JsonRes(w, result, http.StatusOK)
}

func (c *Users) Destroy(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")

	err := c.usecase.Destroy(id)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, c.configs.Exceptions.ErrParseId) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	output["message"] = c.configs.Translations.Users.Destroy.Success
	response.JsonRes(w, output, http.StatusOK)
}

func (c *Users) Block(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")
	err := c.usecase.BlockUnblock(id, true)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, c.configs.Exceptions.ErrParseId) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	output["message"] = c.configs.Translations.Users.Block.Success
	response.JsonRes(w, output, http.StatusOK)
}

func (c *Users) Unblock(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")
	err := c.usecase.BlockUnblock(id, false)
	if err != nil {
		var status int
		if errors.Is(err, c.configs.Exceptions.ErrUserNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, c.configs.Exceptions.ErrParseId) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	output["message"] = c.configs.Translations.Users.Unblock.Success
	response.JsonRes(w, output, http.StatusOK)
}
