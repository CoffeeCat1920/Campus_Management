package modals

import (
	"time"

	"github.com/google/uuid"
)

type Borrow struct {
	UUID       uuid.UUID `json:"uuid"`
	BookId     string    `json:"bookid"`
	UserId     string    `json:"userid"`
	BorrowDate string    `json:"borrow_date"`
	ReturnDate string    `json:"return_date"`
}

func NewBorrow(bookid string, userid string, days int) *Borrow {
	now := time.Now()
	returnDate := now.AddDate(0, 0, days)

	return &Borrow{
		UUID:       uuid.New(),
		BookId:     bookid,
		UserId:     userid,
		BorrowDate: now.Format("2006 - 01 - 02"),
		ReturnDate: returnDate.Format("2006 - 01 - 02"),
	}
}
