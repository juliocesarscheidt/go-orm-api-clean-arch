package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/repository"
	service "github.com/juliocesarscheidt/go-orm-api/application/service"
)

type CountUsersUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService service.PasswordService
}

func NewCountUsersUsecase(userRepository repository.UserRepository) *CountUsersUsecase {
	return &CountUsersUsecase{
		UserRepository: userRepository,
	}
}

func (usecase *CountUsersUsecase) Execute(countUsersDto *dto.CountUsersDto) (int, error) {
	return usecase.UserRepository.CountUsers()
}
