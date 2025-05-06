package database

import (
	"database/sql"
	"fmt"
	"time"
	modals "what/internal/models"
)

func (s *service) AddBorrow(borrow *modals.Borrow) error {
	q := `
	INSERT INTO borrows(uuid, bookid, userid, borrow_date, return_date)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.db.Exec(q, &borrow.UUID, &borrow.BookId, &borrow.UserId, &borrow.BorrowDate, &borrow.ReturnDate)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBorrow(uuid string) error {
	q := "DELETE FROM borrows WHERE uuid = $1"

	res, err := s.db.Exec(q, uuid)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrItemNotFound
	}

	return nil
}

func (s *service) BorrowFine(uuid string, rate int) (int, error) {
	q := `
	SELECT return_date 
	FROM borrows 
	WHERE uuid = $1 
	`
	var return_date string
	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&return_date)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		} else {
			return 0, err
		}
	}

	layout := "2006-01-02"
	targetDate, err := time.Parse(layout, return_date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return 0, err
	}

	now := time.Now()
	if now.After(targetDate) {
		duration := now.Sub(targetDate)
		days := int(duration.Hours() / 24)
		return days * rate, nil
	} else {
		return 0, nil
	}
}
