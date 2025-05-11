package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) GetAllLibrarians() ([]modals.User, error) {
	var users []modals.User

	q := `SELECT uuid, name, password, type, rentedbooks FROM users WHERE type = 1;`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user modals.User
		err := rows.Scan(&user.UUID, &user.Name, &user.Password, &user.Type, &user.RentedBooks)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) GetLibrarianUUIDFromName(name string) (string, error) {
	var librarianUUID string

	query := "SELECT uuid FROM users WHERE name = $1 AND type = 1"
	err := s.db.QueryRow(query, name).Scan(&librarianUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrItemNotFound
		} else {
			return "", err
		}
	}

	return librarianUUID, nil
}

func (s *service) GetLibrarianFromUUID(uuid string) (*modals.User, error) {
	var librarian modals.User

	query := "SELECT uuid, name, password, type, rentedbooks FROM users WHERE uuid = $1 AND type = 1"
	err := s.db.QueryRow(query, uuid).Scan(&librarian.UUID, &librarian.Name, &librarian.Password, &librarian.Type, &librarian.RentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &librarian, nil
}

func (s *service) GetLibrarianFromName(name string) (*modals.User, error) {
	var librarian modals.User

	query := "SELECT uuid, name, password, type, rentedbooks FROM users WHERE name = $1 AND type = 1"
	err := s.db.QueryRow(query, name).Scan(&librarian.UUID, &librarian.Name, &librarian.Password, &librarian.Type, &librarian.RentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &librarian, nil
}
