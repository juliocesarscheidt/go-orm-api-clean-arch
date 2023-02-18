package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	infrapresenter "github.com/juliocesarscheidt/go-orm-api/infra/presenter"
	infraservice "github.com/juliocesarscheidt/go-orm-api/infra/service"
)

// using pseudo heritage to BaseController
type UserController struct {
	BaseController
	CreateUserUsecase *usecase.CreateUserUsecase
	GetUserUsecase    *usecase.GetUserUsecase
	GetUsersUsecase   *usecase.GetUsersUsecase
	UpdateUserUsecase *usecase.UpdateUserUsecase
	DeleteUserUsecase *usecase.DeleteUserUsecase
	CountUsersUsecase *usecase.CountUsersUsecase
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	passwordService := &infraservice.PasswordService{}
	userPresenter := &infrapresenter.UserPresenter{}
	return &UserController{
		CreateUserUsecase: usecase.NewCreateUserUsecase(userRepository, passwordService),
		GetUserUsecase:    usecase.NewGetUserUsecase(userRepository, userPresenter),
		GetUsersUsecase:   usecase.NewGetUsersUsecase(userRepository, userPresenter),
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
			controller.HandleError(w, r, err)
			return
		}

		id, err := controller.CreateUserUsecase.Execute(createUserDto)
		if err != nil {
			controller.HandleError(w, r, err)
			return
		}

		controller.SendCreated(w, r, id, nil)
	}
}

func (controller UserController) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := controller.GetValueFromParamsAsInt(mux.Vars(r), "id")
		getUserDto := &dto.GetUserDto{Id: int(id)}

		user, err := controller.GetUserUsecase.Execute(getUserDto)
		if err != nil {
			controller.HandleError(w, r, err)
			return
		}
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		controller.SendOk(w, r, user, nil)
	}
}

func (controller UserController) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := controller.GetValueFromFormAsInt(r.FormValue, "page")
		size, _ := controller.GetValueFromFormAsInt(r.FormValue, "size")
		getUsersDto := &dto.GetUsersDto{Page: page, Size: size}

		users, err := controller.GetUsersUsecase.Execute(getUsersDto)
		if err != nil {
			controller.HandleError(w, r, err)
			return
		}

		counter, err := controller.CountUsersUsecase.Execute(&dto.CountUsersDto{})
		if err != nil {
			controller.HandleError(w, r, err)
			return
		}

		controller.SendOk(w, r, users, map[string]int{"total": counter})
	}
}

func (controller UserController) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var updateUserDto *dto.UpdateUserDto
		if err := json.NewDecoder(r.Body).Decode(&updateUserDto); err != nil {
			controller.HandleError(w, r, err)
			return
		}
		id, _ := controller.GetValueFromParamsAsInt(mux.Vars(r), "id")
		updateUserDto.Id = int(id)
		if err := controller.UpdateUserUsecase.Execute(updateUserDto); err != nil {
			controller.HandleError(w, r, err)
			return
		}

		controller.SendAccepted(w, r, "Success")
	}
}

func (controller UserController) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := controller.GetValueFromParamsAsInt(mux.Vars(r), "id")
		deleteUserDto := &dto.DeleteUserDto{Id: id}
		if err := controller.DeleteUserUsecase.Execute(deleteUserDto); err != nil {
			controller.HandleError(w, r, err)
			return
		}

		controller.SendNoContent(w, r)
	}
}
