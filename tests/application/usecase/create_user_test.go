package usecase

import (
	"errors"
	"testing"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	infraservice "github.com/juliocesarscheidt/go-orm-api/infra/service"
)

func TestCreateUserSuccess(t *testing.T) {
	passwordService := &infraservice.PasswordService{}
	userRepository := repository.UserRepositoryMemory{}
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	getUserUsecase := usecase.NewGetUserUsecase(userRepository)
	deleteUserUsecase := usecase.NewDeleteUserUsecase(userRepository)
	// create a user
	createUserDto := &dto.CreateUserDto{
		Name:     "test",
		Email:    "test@mail.com",
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
	getUserDto := &dto.GetUserDto{Id: id}
	user, err := getUserUsecase.Execute(getUserDto)
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
	deleteUserDto := &dto.DeleteUserDto{Id: id}
	if err := deleteUserUsecase.Execute(deleteUserDto); err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestCreateUserInvalidPassword(t *testing.T) {
	passwordService := &infraservice.PasswordService{}
	userRepository := repository.UserRepositoryMemory{}
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	// create a user
	createUserDto := &dto.CreateUserDto{
		Name:     "test",
		Email:    "test@mail.com",
		Password: "a",
	}
	expectedErr := errors.New("Invalid password length")
	_, err := createUserUsecase.Execute(createUserDto)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected err to be %v, got %v", expectedErr, err)
	}
}

// go vet
// go test tests/**/**/*_test.go -v

// go test -cover -coverpkg=github.com/juliocesarscheidt/go-orm-api/application/usecase -coverprofile cover.out tests/**/**/*_test.go -v
// go tool cover -html=cover.out -o coverage.html

// go test tests\application\usecase\create_user_test.go -v
