package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type UpdateUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func (usecase *UpdateUserUsecase) Execute(updateUserDto *dto.UpdateUserDto) error {
	userBuilder := entity.UserBuilder{PasswordService: usecase.PasswordService}
	user, err := userBuilder.AlterUser(updateUserDto.Name, updateUserDto.Password)
	if err != nil {
		return err
	}
	if err := usecase.UserRepository.UpdateUser(updateUserDto.Id, user); err != nil {
		return err
	}
	return nil
}
