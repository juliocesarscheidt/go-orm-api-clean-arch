package usecase

import (
	"testing"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	infraservice "github.com/juliocesarscheidt/go-orm-api/infra/service"
)

func TestGetUsersEmptySuccess(t *testing.T) {
	userRepository := repository.UserRepositoryMemory{}
	getUsersUsecase := usecase.NewGetUsersUsecase(userRepository)
	getUsersDto := &dto.GetUsersDto{Page: 0, Size: 1}
	users, err := getUsersUsecase.Execute(getUsersDto)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if users != nil {
		t.Errorf("Expected user to be nil, got %v", users)
	}
}

func TestGetUsersNonEmptySuccess(t *testing.T) {
	passwordService := &infraservice.PasswordService{}
	userRepository := repository.UserRepositoryMemory{}
	// create some user
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	createUserDto := &dto.CreateUserDto{
		Name:     "test",
		Email:    "test@mail.com",
		Password: "PASSWORD",
	}
	id, err := createUserUsecase.Execute(createUserDto)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	// retrieve users of page 0 (it should return the user)
	getUsersUsecase := usecase.NewGetUsersUsecase(userRepository)
	users, err := getUsersUsecase.Execute(&dto.GetUsersDto{Page: 0, Size: 10})
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if users == nil {
		t.Errorf("Expected user to be not nil, got %v", users)
	}
	if len(users) != 1 {
		t.Errorf("Expected user to have length of 1, got %v", len(users))
	}
	// check the count of users
	countUsersUsecase := usecase.NewCountUsersUsecase(userRepository)
	counter, err := countUsersUsecase.Execute(&dto.CountUsersDto{})
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if counter != 1 {
		t.Errorf("Expected counter to be 1, got %v", counter)
	}
	// retrieve users of page 1 (it should not return anything)
	users, _ = getUsersUsecase.Execute(&dto.GetUsersDto{Page: 1, Size: 10})
	if users != nil {
		t.Errorf("Expected user to be not nil, got %v", users)
	}
	if len(users) != 0 {
		t.Errorf("Expected user to have length of 0, got %v", len(users))
	}
	// remove created user
	deleteUserUsecase := usecase.NewDeleteUserUsecase(userRepository)
	if err := deleteUserUsecase.Execute(&dto.DeleteUserDto{Id: id}); err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

// go vet
// go test tests/**/**/*_test.go -v

// go test -cover -coverpkg=github.com/juliocesarscheidt/go-orm-api/application/usecase -coverprofile cover.out tests/**/**/*_test.go -v
// go tool cover -html=cover.out -o coverage.html

// go test tests\application\usecase\get_users_test.go -v
