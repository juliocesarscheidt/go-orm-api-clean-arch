package presenter

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
)

type UserPresenter struct {
}

func (presenter *UserPresenter) Map(user *entity.User) *dto.UserViewDto {
	return &dto.UserViewDto{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: utils.FormatTimeToRFC3339(user.CreatedAt),
		UpdatedAt: utils.FormatTimeToRFC3339(user.UpdatedAt),
		DeletedAt: utils.FormatTimeToRFC3339(user.DeletedAt.Time),
	}
}

func (presenter *UserPresenter) MapCollection(users []*entity.User) []*dto.UserViewDto {
	var usersDto []*dto.UserViewDto
	for _, user := range users {
		usersDto = append(usersDto, presenter.Map(user))
	}
	return usersDto
}
