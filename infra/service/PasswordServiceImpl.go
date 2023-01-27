package service

import (
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
	"golang.org/x/crypto/bcrypt"
)

// using pseudo polymorphism
type PasswordService struct {
}

func (passwordService *PasswordService) EncryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	utils.Logger.Infof("Password :: %v", string(password))
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	utils.Logger.Infof("Hashed Password :: %v", string(hashedPassword))
	return string(hashedPassword), nil
}

func (passwordService *PasswordService) ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
