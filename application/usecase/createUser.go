package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type CreateUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func (usecase *CreateUserUsecase) Execute(createUserDto *dto.CreateUserDto) error {
	userBuilder := entity.UserBuilder{PasswordService: usecase.PasswordService}
	user, err := userBuilder.NewUser(createUserDto.Name, createUserDto.Email, createUserDto.Password)
	if err != nil {
		return err
	}
	usecase.UserRepository.CreateUser(user)
	return nil
}
