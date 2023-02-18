package usecase

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	applicationpresenter "github.com/juliocesarscheidt/go-orm-api/application/presenter"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
)

type GetUserUsecase struct {
	UserRepository repository.UserRepository
	UserPresenter  applicationpresenter.UserPresenter
}

func NewGetUserUsecase(userRepository repository.UserRepository, userPresenter applicationpresenter.UserPresenter) *GetUserUsecase {
	return &GetUserUsecase{
		UserRepository: userRepository,
		UserPresenter:  userPresenter,
	}
}

func (usecase *GetUserUsecase) Execute(getUserDto *dto.GetUserDto) (*dto.UserViewDto, error) {
	user, err := usecase.UserRepository.GetUser(getUserDto.Id)
	if err != nil {
		return nil, err
	}
	return usecase.UserPresenter.Map(user), nil
}
