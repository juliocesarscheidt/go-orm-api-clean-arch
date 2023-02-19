package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	applicationpresenter "github.com/juliocesarscheidt/go-orm-api/application/presenter"
	"github.com/juliocesarscheidt/go-orm-api/application/repository"
)

type GetUsersUsecase struct {
	UserRepository repository.UserRepository
	UserPresenter  applicationpresenter.UserPresenter
}

func NewGetUsersUsecase(userRepository repository.UserRepository, userPresenter applicationpresenter.UserPresenter) *GetUsersUsecase {
	return &GetUsersUsecase{
		UserRepository: userRepository,
		UserPresenter:  userPresenter,
	}
}

func (usecase *GetUsersUsecase) Execute(getUsersDto *dto.GetUsersDto) ([]*dto.UserViewDto, error) {
	users, err := usecase.UserRepository.GetUsers(getUsersDto.Page, getUsersDto.Size)
	if err != nil {
		return nil, err
	}
	return usecase.UserPresenter.MapCollection(users), nil
}
