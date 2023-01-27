package dto

import (
	"time"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (createUserDto *CreateUserDto) NewUser(passwordService service.PasswordService) *entity.User {
	password, _ := passwordService.EncryptPassword(createUserDto.Password)
	return &entity.User{
		Name:      createUserDto.Name,
		Email:     createUserDto.Email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
