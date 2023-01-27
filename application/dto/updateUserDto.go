package dto

import (
	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type UpdateUserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (updateUserDto *UpdateUserDto) NewUser(passwordService service.PasswordService) *entity.User {
	password, _ := passwordService.EncryptPassword(updateUserDto.Password)
	return &entity.User{
		Name:     updateUserDto.Name,
		Password: password,
	}
}
