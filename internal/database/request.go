package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) AddRequest(request *modals.Request) error {
	q := `
	INSERT INTO requests (uuid, userid, bookid)
	VALUES ($1, $2, $3)
	`
	_, err := s.db.Exec(q, request.UUID, request.UserId, request.BookId)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetRequest(uuid string) (*modals.Request, error) {
	request := modals.Request{}
	q := `SELECT * FROM requests WHERE uuid = $1`

	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&request.UUID, &request.UserId, &request.BookId)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (s *service) GetRequestsByUse(uuid string) ([]modals.RequestWithBookName, error) {
	requests := []modals.RequestWithBookName{}

	q := `SELECT uuid, userid, bookid FROM requests WHERE userid = $1`

	rows, err := s.db.Query(q, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return requests, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var request modals.Request
		err := rows.Scan(&request.UUID, &request.UserId, &request.BookId)
		if err != nil {
			return nil, err
		}

		book, err := s.GetBookFromUUID(request.BookId)
		if err != nil {
			if err == sql.ErrNoRows {
				continue // or return nil, ErrItemNotFound depending on your needs
			}
			return nil, err
		}

		requests = append(requests, modals.RequestWithBookName{
			UUID:     request.UUID.String(),
			Bookname: book.Name,
			Userid:   request.UserId,
			Bookid:   request.BookId,
		})
	}

	if len(requests) == 0 {
		return nil, ErrItemNotFound
	}

	return requests, nil
}

func (s *service) DeleteRequest(uuid string) error {
	q := "DELETE FROM requests WHERE uuid = $1"

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
