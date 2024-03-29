package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/repository"
	service "github.com/juliocesarscheidt/go-orm-api/application/service"
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
)

type UpdateUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService service.PasswordService
}

func NewUpdateUserUsecase(userRepository repository.UserRepository, passwordService service.PasswordService) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		UserRepository:  userRepository,
		PasswordService: passwordService,
	}
}

func (usecase *UpdateUserUsecase) Execute(updateUserDto *dto.UpdateUserDto) error {
	err := entity.ValidateUserFields(map[string]string{"name": updateUserDto.Name, "password": updateUserDto.Password})
	if err != nil {
		return err
	}
	user, err := usecase.UserRepository.GetUser(updateUserDto.Id)
	if err != nil {
		return err
	}
	hashedPassword, _ := usecase.PasswordService.EncryptPassword(updateUserDto.Password)
	user.UpdateUser(updateUserDto.Name, hashedPassword)
	return usecase.UserRepository.UpdateUser(updateUserDto.Id, user)
}
