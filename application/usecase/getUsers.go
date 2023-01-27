package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	applicationservice "github.com/juliocesarscheidt/go-orm-api/application/service"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
)

type GetUsersUsecase struct {
	UserRepository repository.UserRepository
}

func (usecase *GetUsersUsecase) Execute(getUsersDto *dto.GetUsersDto) ([]*dto.UserViewDto, error) {
	users, err := usecase.UserRepository.GetUsers(getUsersDto.Page, getUsersDto.Size)
	if err != nil {
		return nil, err
	}
	var usersDto []*dto.UserViewDto
	for _, user := range users {
		usersDto = append(usersDto, applicationservice.MapUserToDto(user))
	}
	return usersDto, nil
}
