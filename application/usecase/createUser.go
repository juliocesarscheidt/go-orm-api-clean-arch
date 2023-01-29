package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/builder"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type CreateUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func (usecase *CreateUserUsecase) Execute(createUserDto *dto.CreateUserDto) (int, error) {
	userBuilder := builder.UserBuilder{PasswordService: usecase.PasswordService}
	user, err := userBuilder.NewUser(createUserDto.Name, createUserDto.Email, createUserDto.Password)
	if err != nil {
		return 0, err
	}
	return usecase.UserRepository.CreateUser(user)
}
