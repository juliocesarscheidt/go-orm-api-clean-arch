package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type CountUsersUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func NewCountUsersUsecase(userRepository repository.UserRepository) *CountUsersUsecase {
	return &CountUsersUsecase{
		UserRepository: userRepository,
	}
}

func (usecase *CountUsersUsecase) Execute(countUsersDto *dto.CountUsersDto) (int, error) {
	return usecase.UserRepository.CountUsers()
}
