package presenter

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
)

type UserPresenter interface {
	Map(user *entity.User) *dto.UserViewDto
	MapCollection(users []*entity.User) []*dto.UserViewDto
}
