package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
)

func GetUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")

		getUserDto := &dto.GetUserDto{Id: int(id)}

		getUserUsecase := usecase.GetUserUsecase{UserRepository: userRepository}
		user, err := getUserUsecase.Execute(getUserDto)
		if err != nil {
			HandleError(w, err)
			return
		}
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&dto.HttpResponseDto{Data: user})
	}
}

func GetUsers(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := GetValueFromFormAsInt(r.FormValue, "page")
		size, _ := GetValueFromFormAsInt(r.FormValue, "size")
		getUsersDto := &dto.GetUsersDto{Page: page, Size: size}

		getUsersUsecase := usecase.GetUsersUsecase{UserRepository: userRepository}
		users, err := getUsersUsecase.Execute(getUsersDto)
		if err != nil {
			HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&dto.HttpResponseDto{Data: users})
	}
}

func CreateUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createUserDto *dto.CreateUserDto
		if err := json.NewDecoder(r.Body).Decode(&createUserDto); err != nil {
			HandleError(w, err)
			return
		}

		createUserUsecase := usecase.CreateUserUsecase{UserRepository: userRepository}
		err := createUserUsecase.Execute(createUserDto)
		if err != nil {
			HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: "Success"})
	}
}

func UpdateUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var updateUserDto *dto.UpdateUserDto
		if err := json.NewDecoder(r.Body).Decode(&updateUserDto); err != nil {
			HandleError(w, err)
			return
		}
		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		updateUserDto.Id = int(id)

		updateUserUsecase := usecase.UpdateUserUsecase{UserRepository: userRepository}
		if err := updateUserUsecase.Execute(updateUserDto); err != nil {
			HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: "Success"})
	}
}

func DeleteUser(userRepository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := GetValueFromParamsAsInt(mux.Vars(r), "id")
		deleteUserDto := &dto.DeleteUserDto{Id: id}

		deleteUserUsecase := usecase.DeleteUserUsecase{UserRepository: userRepository}
		if err := deleteUserUsecase.Execute(deleteUserDto); err != nil {
			HandleError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
