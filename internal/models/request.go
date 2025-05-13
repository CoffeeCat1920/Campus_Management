package modals

import "github.com/google/uuid"

// Only to be used once, cause of a WEIRD db err on Linux
// TODO: Add the book name to the actual Request struct
type RequestWithBookName struct {
	UUID     string `json:"uuid"`
	Bookname string `json:"bookname"`
	Userid   string `json:"userid"`
	Bookid   string `json:"bookid"`
}

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
