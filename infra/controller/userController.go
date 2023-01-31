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

var passwordService *infraservice.PasswordService

func init() {
	passwordService = &infraservice.PasswordService{}
}

func GetUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		getUserDto := &dto.GetUserDto{Id: int(id)}

		getUserUsecase := usecase.GetUserUsecase{UserRepository: userRepository, PasswordService: passwordService}
		user, err := getUserUsecase.Execute(getUserDto)
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

func GetUsers(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := GetValueFromFormAsInt(r.FormValue, "page")
		size, _ := GetValueFromFormAsInt(r.FormValue, "size")
		getUsersDto := &dto.GetUsersDto{Page: page, Size: size}

		getUsersUsecase := usecase.GetUsersUsecase{UserRepository: userRepository, PasswordService: passwordService}
		users, err := getUsersUsecase.Execute(getUsersDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}

		countUsersDto := &dto.CountUsersDto{}
		countUsersUsecase := usecase.CountUsersUsecase{UserRepository: userRepository, PasswordService: passwordService}
		counter, err := countUsersUsecase.Execute(countUsersDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}

		SendOk(w, r, users, map[string]int{"total": counter})
	}
}

func CreateUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createUserDto *dto.CreateUserDto
		if err := json.NewDecoder(r.Body).Decode(&createUserDto); err != nil {
			HandleError(w, r, err)
			return
		}

		createUserUsecase := usecase.CreateUserUsecase{UserRepository: userRepository, PasswordService: passwordService}
		id, err := createUserUsecase.Execute(createUserDto)
		if err != nil {
			HandleError(w, r, err)
			return
		}

		SendCreated(w, r, id, nil)
	}
}

func UpdateUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var updateUserDto *dto.UpdateUserDto
		if err := json.NewDecoder(r.Body).Decode(&updateUserDto); err != nil {
			HandleError(w, r, err)
			return
		}
		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		updateUserDto.Id = int(id)

		updateUserUsecase := usecase.UpdateUserUsecase{UserRepository: userRepository, PasswordService: passwordService}
		if err := updateUserUsecase.Execute(updateUserDto); err != nil {
			HandleError(w, r, err)
			return
		}

		SendAccepted(w, r, "Success")
	}
}

func DeleteUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		deleteUserDto := &dto.DeleteUserDto{Id: id}

		deleteUserUsecase := usecase.DeleteUserUsecase{UserRepository: userRepository, PasswordService: passwordService}
		if err := deleteUserUsecase.Execute(deleteUserDto); err != nil {
			HandleError(w, r, err)
			return
		}

		SendNoContent(w, r)
	}
}
