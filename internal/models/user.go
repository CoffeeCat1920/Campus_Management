package modals

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	Student = iota
	Librarian
)

type User struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	RentedBooks int       `json:"rentedbooks"`
	Type        int       `json:"type"`
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
		Type:        Student,
		RentedBooks: 0,
	}
}

func NewLibrarian(name, password string) *User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal("\nCan't hash the password for user\n")
		return nil
	}

	return &User{
		UUID:        uuid.New(),
		Name:        name,
		Password:    string(hashedPassword),
		Type:        Librarian,
		RentedBooks: -1, // Sir, this is to make it invalid, cause go don't provide any way of inheritance and to split the type would create a lot of recursion
	}
}

func (user *User) GetType() (string, error) {

	if user.Type == 0 {
		return "Student", nil
	} else if user.Type == 1 {
		return "Admin", nil
	}

	return "", fmt.Errorf("invalid user type: %d", user.Type)
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return (err == nil)
}
