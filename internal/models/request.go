package modals

import "github.com/google/uuid"

type Request struct {
	UUID   uuid.UUID `json:"uuid"`
	UserId string    `json:"userid"`
	BookId string    `json:"bookid"`
}

func NewRequest(userId string, bookId string) *Request {
	return &Request{
		UUID:   uuid.New(),
		UserId: userId,
		BookId: bookId,
	}
}
