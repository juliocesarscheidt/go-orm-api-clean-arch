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

// go vet
// go test tests/**/**/*_test.go -v

// go test -cover -coverpkg=github.com/juliocesarscheidt/go-orm-api/application/usecase -coverprofile cover.out tests/**/**/*_test.go -v
// go tool cover -html=cover.out -o coverage.html

// go test tests\application\usecase\get_user_test.go -v
