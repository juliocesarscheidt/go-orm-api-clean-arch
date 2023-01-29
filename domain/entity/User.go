package entity

import (
	"errors"
	"fmt"
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
			errorMessages = append(errorMessages, fmt.Sprintf("Invalid %s", strings.Title(key)))
		}
	}
	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, ", "))
	}
	return nil
}
