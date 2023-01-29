package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	applicationservice "github.com/juliocesarscheidt/go-orm-api/application/service"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type GetUserUsecase struct {
	UserRepository  repository.UserRepository
	PasswordService domainservice.PasswordService
}

func (usecase *GetUserUsecase) Execute(getUserDto *dto.GetUserDto) (*dto.UserViewDto, error) {
	user, err := usecase.UserRepository.GetUser(getUserDto.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return applicationservice.MapUserToDto(user), nil
}
