package modals

import "github.com/google/uuid"

type Book struct {
	UUID      uuid.UUID `json:"uuid"`
	ISBN      string    `json:"isbn"`
	Name      string    `json:"name"`
	Available bool      `json:"available"`
}

func NewBook(isbn, name string) *Book {
	return &Book{
		UUID:      uuid.New(),
		ISBN:      isbn,
		Name:      name,
		Available: true,
	}
}
