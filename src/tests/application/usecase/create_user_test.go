package usecase

import (
	"errors"
	"testing"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/infra/presenter"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/service"
)

func TestCreateUserSuccess(t *testing.T) {
	passwordService := &service.PasswordService{}
	userPresenter := &presenter.UserPresenter{}
	userRepository := repository.UserRepositoryMemory{}
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	getUserUsecase := usecase.NewGetUserUsecase(userRepository, userPresenter)
	deleteUserUsecase := usecase.NewDeleteUserUsecase(userRepository)
	// create a user
	createUserDto := &dto.CreateUserDto{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "PASSWORD",
	}
	id, err := createUserUsecase.Execute(createUserDto)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if id <= 0 {
		t.Errorf("Expected id to be greater than 0, got %v", id)
	}
	// retrieve the just created user
	user, err := getUserUsecase.Execute(&dto.GetUserDto{Id: id})
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if user == nil {
		t.Errorf("Expected user to exist, got %v", user)
	}
	if user.Email != createUserDto.Email {
		t.Errorf("Expected user email to match the DTO email, got %v", user.Email)
	}
	// remove created user
	if err := deleteUserUsecase.Execute(&dto.DeleteUserDto{Id: id}); err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestCreateUserInvalidPassword(t *testing.T) {
	passwordService := &service.PasswordService{}
	userRepository := repository.UserRepositoryMemory{}
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	// create a user
	createUserDto := &dto.CreateUserDto{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "a",
	}
	expectedErr := errors.New("Invalid password length, the password must have at least 8 and at most 50 characters")
	_, err := createUserUsecase.Execute(createUserDto)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected err to be %v, got %v", expectedErr, err)
	}
}

func TestCreateUserInvalidFields(t *testing.T) {
	passwordService := &service.PasswordService{}
	userRepository := repository.UserRepositoryMemory{}
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	// create a user
	createUserDto := &dto.CreateUserDto{
		Name:     "",
		Email:    "",
		Password: "",
	}
	expectedErr := errors.New("Invalid email, Invalid name, Invalid password")
	_, err := createUserUsecase.Execute(createUserDto)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected err to be %v, got %v", expectedErr, err)
	}
}
