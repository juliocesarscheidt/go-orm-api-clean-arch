package service

import (
	"golang.org/x/crypto/bcrypt"
)

// using pseudo polymorphism
type PasswordService struct {
}

func (passwordService *PasswordService) EncryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (passwordService *PasswordService) ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
