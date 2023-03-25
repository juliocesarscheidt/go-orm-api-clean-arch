package usecase

import (
	"testing"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/infra/presenter"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/service"
)

func TestListUsersEmptySuccess(t *testing.T) {
	userRepository := repository.UserRepositoryMemory{}
	userPresenter := &presenter.UserPresenter{}
	listUsersUsecase := usecase.NewListUsersUsecase(userRepository, userPresenter)
	listUsersDto := &dto.ListUsersDto{Page: 0, Size: 1}
	users, err := listUsersUsecase.Execute(listUsersDto)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if users != nil {
		t.Errorf("Expected user to be nil, got %v", users)
	}
}

func TestListUsersNonEmptySuccess(t *testing.T) {
	passwordService := &service.PasswordService{}
	userPresenter := &presenter.UserPresenter{}
	userRepository := repository.UserRepositoryMemory{}
	// create some user
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	createUserDto := &dto.CreateUserDto{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "PASSWORD",
	}
	id, err := createUserUsecase.Execute(createUserDto)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	// retrieve users of page 0 (it should return the user)
	listUsersUsecase := usecase.NewListUsersUsecase(userRepository, userPresenter)
	users, err := listUsersUsecase.Execute(&dto.ListUsersDto{Page: 0, Size: 10})
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
	users, _ = listUsersUsecase.Execute(&dto.ListUsersDto{Page: 1, Size: 10})
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
