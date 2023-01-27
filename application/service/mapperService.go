package service

import (
	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
)

func MapUserToDto(user *entity.User) *dto.UserViewDto {
	return &dto.UserViewDto{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: FormatTimeToRFC(user.CreatedAt),
		UpdatedAt: FormatTimeToRFC(user.UpdatedAt),
		DeletedAt: FormatTimeToRFC(user.DeletedAt.Time),
	}
}
