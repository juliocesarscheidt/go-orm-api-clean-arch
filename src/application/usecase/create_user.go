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

func NewCreateUserUsecase(userRepository repository.UserRepository, passwordService domainservice.PasswordService) *CreateUserUsecase {
	return &CreateUserUsecase{
		UserRepository:  userRepository,
		PasswordService: passwordService,
	}
}

func (usecase *CreateUserUsecase) Execute(createUserDto *dto.CreateUserDto) (int, error) {
	err := entity.ValidateUserFields(map[string]string{"name": createUserDto.Name, "email": createUserDto.Email, "password": createUserDto.Password})
	if err != nil {
		return 0, err
	}
	hashedPassword, _ := usecase.PasswordService.EncryptPassword(createUserDto.Password)
	user := entity.NewUser(createUserDto.Name, createUserDto.Email, hashedPassword)
	return usecase.UserRepository.CreateUser(user)
}
