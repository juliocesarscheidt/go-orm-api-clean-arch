package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/repository"
	service "github.com/juliocesarscheidt/go-orm-api/application/service"
)

type DeleteUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService service.PasswordService
}

func NewDeleteUserUsecase(userRepository repository.UserRepository) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		UserRepository: userRepository,
	}
}

func (usecase *DeleteUserUsecase) Execute(deleteUserDto *dto.DeleteUserDto) error {
	return usecase.UserRepository.DeleteUser(deleteUserDto.Id)
}
