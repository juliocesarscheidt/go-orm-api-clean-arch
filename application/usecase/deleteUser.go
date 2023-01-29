package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type DeleteUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func (usecase *DeleteUserUsecase) Execute(deleteUserDto *dto.DeleteUserDto) error {
	return usecase.UserRepository.DeleteUser(deleteUserDto.Id)
}
