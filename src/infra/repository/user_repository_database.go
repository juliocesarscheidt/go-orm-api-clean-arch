package repository

import (
	"errors"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryDatabase struct {
	Db *gorm.DB
}

func (userRepository UserRepositoryDatabase) MigrateUser() error {
	return userRepository.Db.AutoMigrate(&entity.User{})
}

func (userRepository UserRepositoryDatabase) ListUsers(page, size int) ([]*entity.User, error) {
	var users []*entity.User
	userRepository.Db.Limit(size).Offset(page * size).Find(&users)
	if len(users) == 0 {
		return nil, nil
	}
	return users, nil
}

func (userRepository UserRepositoryDatabase) GetUser(id int) (*entity.User, error) {
	var user *entity.User
	result := userRepository.Db.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("Not found")
	}
	if user == nil {
		return nil, errors.New("Not found")
	}
	return user, nil
}

func (userRepository UserRepositoryDatabase) CreateUser(user *entity.User) (int, error) {
	result := userRepository.Db.Create(user)
	if result.RowsAffected == 0 {
		return 0, errors.New("Internal server error")
	}
	return user.Id, nil
}

func (userRepository UserRepositoryDatabase) UpdateUser(id int, user *entity.User) error {
	result := userRepository.Db.Model(&entity.User{Id: id}).Updates(user)
	if result.RowsAffected == 0 {
		return errors.New("Not found")
	}
	return nil
}

func (userRepository UserRepositoryDatabase) DeleteUser(id int) error {
	result := userRepository.Db.Delete(&entity.User{Id: id})
	if result.RowsAffected == 0 {
		return errors.New("Not found")
	}
	return nil
}

func (userRepository UserRepositoryDatabase) CountUsers() (int, error) {
	// using raw query
	var counter int
	result := userRepository.Db.Raw("SELECT COUNT(id) as counter FROM `users` WHERE `users`.`deleted_at` IS NULL").Scan(&counter)
	if result.RowsAffected == 0 {
		return 0, errors.New("Internal server error")
	}
	return counter, nil
}
