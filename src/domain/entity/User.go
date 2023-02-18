package entity

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

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

func ValidateUserFields(fields map[string]string) error {
	errorMessages := []string{}
	for key, field := range fields {
		if field == "" {
			errorMessages = append(errorMessages, fmt.Sprintf("Invalid %s", strings.ToLower(key)))
		}
	}
	if len(errorMessages) > 0 {
		sort.Strings(errorMessages)
		return errors.New(strings.Join(errorMessages, ", "))
	}
	if len(fields["password"]) < 8 || len(fields["password"]) > 50 {
		return errors.New("Invalid password length")
	}
	return nil
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (user *User) UpdateUser(name, password string) {
	user.Name = name
	user.Password = password
	user.UpdatedAt = time.Now()
}
