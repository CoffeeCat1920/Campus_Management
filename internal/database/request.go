package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) AddRequest(request *modals.Request) error {
	q := `
	INSERT INTO requests (uuid, userid, bookid, days)
	VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.Exec(q, request.UUID, request.UserId, request.BookId, request.Days)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetRequest(uuid string) (*modals.Request, error) {
	var request modals.Request
	q := `SELECT * FROM requests 
	WHERE uuid = $1`

	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&request.UUID, &request.UserId, &request.BookId, &request.Days)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &request, err
}

func (s *service) DeleteRequest(uuid string) error {
	q := "DELETE FROM books WHERE uuid = $1"

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
