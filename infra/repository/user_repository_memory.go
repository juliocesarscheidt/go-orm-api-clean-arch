package repository

import (
	"errors"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
)

type UserRepositoryMemory struct {
}

var users []*entity.User
var lastUserId = 0

func (userRepository UserRepositoryMemory) GetUsers(page, size int) ([]*entity.User, error) {
	startIndex := page * size
	endIndex := page*size + size

	if startIndex >= len(users) {
		return nil, nil

	} else {
		if len(users) < endIndex && startIndex == 0 {
			return users, nil

		} else if size > len(users) || len(users) < endIndex {
			usersSlice := users[startIndex:]
			return usersSlice, nil
		}
	}

	usersSlice := users[startIndex:endIndex]
	return usersSlice, nil
}

func (userRepository UserRepositoryMemory) GetUser(id int) (*entity.User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("Not found")
}

func (userRepository UserRepositoryMemory) CreateUser(user *entity.User) (int, error) {
	lastUserId = lastUserId + 1
	user.Id = lastUserId
	users = append(users, user)
	return user.Id, nil
}

func (userRepository UserRepositoryMemory) UpdateUser(id int, user *entity.User) error {
	for idx, u := range users {
		if u.Id == id {
			users[idx].Name = user.Name
			users[idx].Password = user.Password
			return nil
		}
	}
	return errors.New("Not found")
}

func (userRepository UserRepositoryMemory) DeleteUser(id int) error {
	for idx, u := range users {
		if u.Id == id {
			// remove from slice
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}

func (userRepository UserRepositoryMemory) CountUsers() (int, error) {
	return len(users), nil
}
