package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	infraservice "github.com/juliocesarscheidt/go-orm-api/infra/service"
)

type CreateUserUsecase struct {
	UserRepository repository.UserRepository
}

func (usecase *CreateUserUsecase) Execute(createUserDto *dto.CreateUserDto) error {
	passwordService := &infraservice.PasswordService{}
	user := createUserDto.NewUser(passwordService)
	usecase.UserRepository.CreateUser(user)
	return nil
}
