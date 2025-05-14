package database

import (
	"database/sql"
	"fmt"
	"time"
	modals "what/internal/models"
)

func (s *service) AddBorrow(borrow *modals.Borrow) error {
	q := `
	INSERT INTO borrows(uuid, bookid, userid, borrow_date, return_date, returned)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := s.db.Exec(q, &borrow.UUID, &borrow.BookId, &borrow.UserId, &borrow.BorrowDate, &borrow.ReturnDate, borrow.Returned)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetBorrowFromUUId(uuid string) (*modals.Borrow, error) {
	var borrow modals.Borrow

	q := `SELECT * FROM borrows 
	WHERE uuid = $1`

	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&borrow.UUID, &borrow.BookId, &borrow.UserId, &borrow.BorrowDate, &borrow.ReturnDate, &borrow.Returned)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &borrow, err
}

func (s *service) ReturnBorrow(uuid string) error {
	q := `
	UPDATE borrows
	SET returned = TRUE
	WHERE uuid = $1
	`

	res, err := s.db.Exec(q, uuid)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	return nil
}

func (s *service) ClearFine(userid string) error {
	q := "DELETE FROM borrows WHERE userid = $1"

	res, err := s.db.Exec(q, userid)
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
	SELECT return_date, returned 
	FROM borrows 
	WHERE uuid = $1 
	`

	var return_date string
	var returned bool
	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&return_date, &returned)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		} else {
			return 0, err
		}
	}

	if returned {
		return 0, nil
	}

	layout := "2006 - 01 - 02"
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

func (s *service) GetBorrowsByUser(uuid string) ([]modals.Borrow, error) {
	borrows := []modals.Borrow{}

	q := `SELECT * FROM borrows WHERE userid = $1 AND returned = false`

	rows, err := s.db.Query(q, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return borrows, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var borrow modals.Borrow
		err := rows.Scan(&borrow.UUID, &borrow.BookId, &borrow.UserId, &borrow.BorrowDate, &borrow.ReturnDate, &borrow.Returned)
		if err != nil {
			return nil, err
		}
		borrows = append(borrows, borrow)
	}

	return borrows, nil
}

func (s *service) GetBorrowsByUserWithBookName(uuid string) ([]modals.BorrowWithBookName, error) {
	borrows := []modals.BorrowWithBookName{}

	q := `SELECT * FROM borrows WHERE userid = $1 AND returned = false`

	rows, err := s.db.Query(q, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return borrows, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var borrow modals.Borrow
		err := rows.Scan(&borrow.UUID, &borrow.BookId, &borrow.UserId, &borrow.BorrowDate, &borrow.ReturnDate, &borrow.Returned)
		if err != nil {
			return nil, err
		}

		book, err := s.GetBookFromUUID(borrow.BookId)
		if err != nil {
			return nil, err
		}

		fine, err := s.BorrowFine(borrow.UUID.String(), 10)

		borrows = append(borrows, modals.BorrowWithBookName{
			UUID:       borrow.UUID,
			BookId:     borrow.BookId,
			UserId:     borrow.UserId,
			BorrowDate: borrow.BorrowDate,
			ReturnDate: borrow.ReturnDate,
			BookName:   book.Name,
			Fine:       fine,
		})
	}

	if len(borrows) == 0 {
		return nil, ErrItemNotFound
	}

	return borrows, nil
}

func (s *service) GetFineByUser(uuid string) (int, error) {
	var fine int

	q := `SELECT * FROM borrows
	WHERE userid = $1`

	rows, err := s.db.Query(q, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var borrow modals.Borrow
		err := rows.Scan(&borrow.UUID, &borrow.BookId, &borrow.UserId, &borrow.BorrowDate, &borrow.ReturnDate, &borrow.Returned)
		if err != nil {
			return 0, err
		}

		f, err := s.BorrowFine(borrow.UUID.String(), 10)
		if err != nil {
			return 0, err
		}

		fine += f
	}

	return fine, nil
}
