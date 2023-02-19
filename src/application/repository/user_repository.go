package repository

import "github.com/juliocesarscheidt/go-orm-api/domain/entity"

type UserRepository interface {
	MigrateUser() error
	GetUsers(page, size int) ([]*entity.User, error)
	GetUser(id int) (*entity.User, error)
	CreateUser(user *entity.User) (int, error)
	UpdateUser(id int, user *entity.User) error
	DeleteUser(id int) error
	CountUsers() (int, error)
}
