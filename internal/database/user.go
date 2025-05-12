package database

import (
	"database/sql"
	modals "what/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) AddUser(student *modals.User) error {
	q := `
  INSERT INTO users(uuid, name, password, type, rentedBooks)
  VALUES($1, $2, $3, $4, $5)
  `
	_, err := s.db.Exec(q, student.UUID, student.Name, student.Password, student.Type, student.RentedBooks)
	if err != nil {
		if IsUniqueViolation(err) {
			return ErrItemAlreadyExists
		} else {
			return err
		}
	}

	return nil
}

func (s *service) UpdateUserFromUUID(uuid string, name string, password string) error {
	q := `
	UPDATE users 
	SET name = $2, password = $3 
	WHERE uuid = $1;
	`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	res, err := s.db.Exec(q, uuid, name, hashedPassword)
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

func (s *service) GetUserFromName(name string) (*modals.User, error) {
	var user modals.User

	query := "SELECT * FROM users WHERE name = $1"
	err := s.db.QueryRow(query, name).Scan(&user.UUID, &user.Name, &user.Password, &user.RentedBooks, &user.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}
