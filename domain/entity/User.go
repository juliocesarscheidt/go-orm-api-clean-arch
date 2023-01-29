package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"

	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UserBuilder struct {
	PasswordService domainservice.PasswordService
}

func (bulder UserBuilder) NewUser(name string, email string, password string) (*User, error) {
	err := validateUserFields(map[string]string{"Name": name, "Email": email, "Password": password})
	if err != nil {
		return nil, err
	}
	hashedPassword, _ := bulder.PasswordService.EncryptPassword(password)
	user := &User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	return user, nil
}

func (bulder UserBuilder) AlterUser(name string, password string) (*User, error) {
	err := validateUserFields(map[string]string{"Name": name, "Password": password})
	if err != nil {
		return nil, err
	}
	hashedPassword, _ := bulder.PasswordService.EncryptPassword(password)
	user := &User{
		Name:      name,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	}
	return user, nil
}

func validateUserFields(fields map[string]string) error {
	for key, field := range fields {
		if field == "" {
			return errors.New(fmt.Sprintf("Invalid %s", strings.Title(key)))
		}
	}
	return nil
}
