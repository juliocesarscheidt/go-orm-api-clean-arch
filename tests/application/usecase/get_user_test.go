package usecase

import (
	"testing"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/usecase"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
)

func TestGetUserSuccess(t *testing.T) {
	userRepository := repository.UserRepositoryMemory{}
	getUserUsecase := usecase.NewGetUserUsecase(userRepository)
	// retrieve an unexisting user
	getUserDto := &dto.GetUserDto{Id: 1}
	user, err := getUserUsecase.Execute(getUserDto)
	if err == nil {
		t.Errorf("Expected err to be not nil, got %v", err)
	}
	if user != nil {
		t.Errorf("Expected user to not exist, got %v", user)
	}
}
