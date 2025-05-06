package modals

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID        uuid.UUID
	Name        string
	Password    string
	RentedBooks int
}

func NewUser(name, password string) *User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal("\nCan't hash the password for user\n")
		return nil
	}

	return &User{
		UUID:        uuid.New(),
		Name:        name,
		Password:    string(hashedPassword),
		RentedBooks: 0,
	}
}
