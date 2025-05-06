package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) AddUser(user *modals.User) error {
	q := `
  INSERT INTO users(uuid, name, password, rentedBooks)
  VALUES($1, $2, $3, $4)
  `
	_, err := s.db.Exec(q, user.UUID, user.Name, user.Password, user.RentedBooks)
	if err != nil {
		if IsUniqueViolation(err) {
			return ErrItemAlreadyExists
		} else {
			return err
		}
	}

	return nil
}

func (s *service) GetUserUUIDFromName(name string) (string, error) {
	var user modals.User

	query := "SELECT * FROM users WHERE name = $1"
	err := s.db.QueryRow(query, name).Scan(&user.UUID, &user.Name, &user.Password, &user.RentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrItemNotFound
		} else {
			return "", err
		}
	}

	return user.UUID.String(), nil
}

func (s *service) NumberOfRentedBooks(uuid string) (int, error) {
	var rentedBooks int

	query := "SELECT rentedbooks FROM users WHERE name = $1"
	err := s.db.QueryRow(query, uuid).Scan(&rentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrItemNotFound
		} else {
			return 0, err
		}
	}

	return rentedBooks, nil
}

func (s *service) IncreaseRented(uuid string) error {
	q := `
    UPDATE users 
    SET rentedbooks = rentedbooks + 1
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

	if rowsAffected <= 0 {
		return ErrItemNotFound
	}

	return nil
}

func (s *service) DecreaseRented(uuid string) error {
	q := `
    UPDATE users 
    SET rentedbooks = rentedbooks - 1
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

	if rowsAffected <= 0 {
		return ErrItemNotFound
	}

	return nil
}
