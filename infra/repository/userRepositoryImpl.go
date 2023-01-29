package repository

import (
	"errors"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (userRepository UserRepository) GetUsers(page, size int) ([]*entity.User, error) {
	var users []*entity.User
	userRepository.Db.Limit(size).Offset(page * size).Find(&users)
	if len(users) == 0 {
		return nil, nil
	}
	return users, nil
}

func (userRepository UserRepository) GetUser(id int) (*entity.User, error) {
	var user *entity.User
	result := userRepository.Db.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return user, nil
}

func (userRepository UserRepository) CreateUser(user *entity.User) (int, error) {
	result := userRepository.Db.Create(user)
	if result.RowsAffected == 0 {
		return 0, errors.New("Internal Server Error")
	}
	return user.Id, nil
}

func (userRepository UserRepository) UpdateUser(id int, user *entity.User) error {
	result := userRepository.Db.Model(&entity.User{Id: id}).Updates(user)
	if result.RowsAffected == 0 {
		return errors.New("Not Found")
	}
	return nil
}

func (userRepository UserRepository) DeleteUser(id int) error {
	result := userRepository.Db.Delete(&entity.User{Id: id})
	if result.RowsAffected == 0 {
		return errors.New("Not Found")
	}
	return nil
}

func (userRepository UserRepository) CountUsers() (int, error) {
	// using raw query
	var counter int
	result := userRepository.Db.Raw("SELECT COUNT(id) as counter FROM `users` WHERE `users`.`deleted_at` IS NULL").Scan(&counter)
	if result.RowsAffected == 0 {
		return 0, errors.New("Not Found")
	}
	return counter, nil
}
