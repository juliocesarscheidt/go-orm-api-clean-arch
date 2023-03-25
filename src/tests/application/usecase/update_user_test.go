package usecase

import (
	"testing"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/infra/presenter"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/service"
)

func TestUpdateUserSuccess(t *testing.T) {
	passwordService := &service.PasswordService{}
	userPresenter := &presenter.UserPresenter{}
	userRepository := repository.UserRepositoryMemory{}
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository, passwordService)
	updateUserUsecase := usecase.NewUpdateUserUsecase(userRepository, passwordService)
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
	// update the user
	updateUserDto := &dto.UpdateUserDto{
		Id:       id,
		Name:     "user changed",
		Password: "PASSWORD_UPDATED",
	}
	if err := updateUserUsecase.Execute(updateUserDto); err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	// retrieve the just created user
	user, err := getUserUsecase.Execute(&dto.GetUserDto{Id: id})
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if user.Name != updateUserDto.Name {
		t.Errorf("Expected user name to match the DTO name, got %v", user.Name)
	}
	// remove created user
	if err := deleteUserUsecase.Execute(&dto.DeleteUserDto{Id: id}); err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}
