package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
)

type DeleteUserUsecase struct {
	UserRepository repository.UserRepository
}

func (usecase *DeleteUserUsecase) Execute(deleteUserDto *dto.DeleteUserDto) error {
	if err := usecase.UserRepository.DeleteUser(deleteUserDto.Id); err != nil {
		return err
	}
	return nil
}
