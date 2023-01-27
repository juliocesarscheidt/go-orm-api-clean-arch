package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	infraservice "github.com/juliocesarscheidt/go-orm-api/infra/service"
)

type UpdateUserUsecase struct {
	UserRepository repository.UserRepository
}

func (usecase *UpdateUserUsecase) Execute(updateUserDto *dto.UpdateUserDto) error {
	passwordService := &infraservice.PasswordService{}
	user := updateUserDto.NewUser(passwordService)
	if err := usecase.UserRepository.UpdateUser(updateUserDto.Id, user); err != nil {
		return err
	}
	return nil
}
