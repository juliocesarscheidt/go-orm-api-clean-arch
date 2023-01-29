package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/builder"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type UpdateUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func (usecase *UpdateUserUsecase) Execute(updateUserDto *dto.UpdateUserDto) error {
	userBuilder := builder.UserBuilder{PasswordService: usecase.PasswordService}
	user, err := userBuilder.AlterUser(updateUserDto.Name, updateUserDto.Password)
	if err != nil {
		return err
	}
	return usecase.UserRepository.UpdateUser(updateUserDto.Id, user)
}
