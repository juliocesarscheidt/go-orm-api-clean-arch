package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	infraservice "github.com/juliocesarscheidt/go-orm-api/infra/service"
)

type UserController struct {
	CreateUserUsecase *usecase.CreateUserUsecase
	GetUserUsecase    *usecase.GetUserUsecase
	GetUsersUsecase   *usecase.GetUsersUsecase
	UpdateUserUsecase *usecase.UpdateUserUsecase
	DeleteUserUsecase *usecase.DeleteUserUsecase
	CountUsersUsecase *usecase.CountUsersUsecase
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	passwordService := &infraservice.PasswordService{}

	return &UserController{
		CreateUserUsecase: usecase.NewCreateUserUsecase(userRepository, passwordService),
		GetUserUsecase:    usecase.NewGetUserUsecase(userRepository),
		GetUsersUsecase:   usecase.NewGetUsersUsecase(userRepository),
		UpdateUserUsecase: usecase.NewUpdateUserUsecase(userRepository, passwordService),
		DeleteUserUsecase: usecase.NewDeleteUserUsecase(userRepository),
		CountUsersUsecase: usecase.NewCountUsersUsecase(userRepository),
	}
}

func (controller UserController) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createUserDto *dto.CreateUserDto
		if err := json.NewDecoder(r.Body).Decode(&createUserDto); err != nil {
			HandleError(w, r, err)
			return
		}

		id, err := controller.CreateUserUsecase.Execute(createUserDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}

		SendCreated(w, r, id, nil)
	}
}

func (controller UserController) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		getUserDto := &dto.GetUserDto{Id: int(id)}

		user, err := controller.GetUserUsecase.Execute(getUserDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		SendOk(w, r, user, nil)
	}
}

func (controller UserController) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := GetValueFromFormAsInt(r.FormValue, "page")
		size, _ := GetValueFromFormAsInt(r.FormValue, "size")
		getUsersDto := &dto.GetUsersDto{Page: page, Size: size}

		users, err := controller.GetUsersUsecase.Execute(getUsersDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}

		countUsersDto := &dto.CountUsersDto{}
		counter, err := controller.CountUsersUsecase.Execute(countUsersDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}

		SendOk(w, r, users, map[string]int{"total": counter})
	}
}

func (controller UserController) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var updateUserDto *dto.UpdateUserDto
		if err := json.NewDecoder(r.Body).Decode(&updateUserDto); err != nil {
			HandleError(w, r, err)
			return
		}
		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		updateUserDto.Id = int(id)
		if err := controller.UpdateUserUsecase.Execute(updateUserDto); err != nil {
			HandleError(w, r, err)
			return
		}

		SendAccepted(w, r, "Success")
	}
}

func (controller UserController) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		deleteUserDto := &dto.DeleteUserDto{Id: id}
		if err := controller.DeleteUserUsecase.Execute(deleteUserDto); err != nil {
			HandleError(w, r, err)
			return
		}

		SendNoContent(w, r)
	}
}
