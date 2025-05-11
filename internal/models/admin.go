package modals

import (
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Password string
}

var (
	admin          *Admin
	admin_password = "123"
)

func NewAdmin() (*Admin, error) {
	if admin != nil {
		return admin, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin_password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	admin = &Admin{
		Password: string(hashedPassword),
	}

	return admin, nil
}

func (admin *Admin) CheckPassword(password string) error {
	hashedPassword := []byte(admin.Password)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)

	return err
}
