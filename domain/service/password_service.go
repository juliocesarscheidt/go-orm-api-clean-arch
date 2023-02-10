package service

type PasswordService interface {
	EncryptPassword(password string) (string, error)
	ComparePasswords(hashedPassword, password string) error
}
