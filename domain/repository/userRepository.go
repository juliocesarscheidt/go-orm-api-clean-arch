package repository

import "github.com/juliocesarscheidt/go-orm-api/domain/entity"

type UserRepository interface {
	GetUsers(page, size int) ([]*entity.User, error)
	GetUser(id int) (*entity.User, error)
	CreateUser(*entity.User) error
	UpdateUser(id int, user *entity.User) error
	DeleteUser(id int) error
}
