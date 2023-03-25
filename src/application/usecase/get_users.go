package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/application/presenter"
	"github.com/juliocesarscheidt/go-orm-api/application/repository"
)

type ListUsersUsecase struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

func NewListUsersUsecase(userRepository repository.UserRepository, userPresenter presenter.UserPresenter) *ListUsersUsecase {
	return &ListUsersUsecase{
		UserRepository: userRepository,
		UserPresenter:  userPresenter,
	}
}

func (usecase *ListUsersUsecase) Execute(listUsersDto *dto.ListUsersDto) ([]*dto.UserViewDto, error) {
	users, err := usecase.UserRepository.ListUsers(listUsersDto.Page, listUsersDto.Size)
	if err != nil {
		return nil, err
	}
	return usecase.UserPresenter.MapCollection(users), nil
}
