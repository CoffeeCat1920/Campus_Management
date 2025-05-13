package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) GetAllStudents() ([]modals.User, error) {
	users := []modals.User{}

	q := `SELECT * FROM users WHERE type = 0;`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user modals.User
		err := rows.Scan(&user.UUID, &user.Name, &user.Password, &user.RentedBooks, &user.Password)
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

func (s *service) GetStudentUUIDFromName(name string) (string, error) {
	var studentName string

	query := "SELECT uuid FROM users WHERE name = $1 AND type = 0"
	err := s.db.QueryRow(query, name).Scan(&studentName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrItemNotFound
		} else {
			return "", err
		}
	}

	return studentName, nil
}

func (s *service) GetStudentFromUUID(uuid string) (*modals.User, error) {
	var student modals.User

	query := "SELECT * FROM users WHERE uuid = $1 AND type = 0"
	err := s.db.QueryRow(query, uuid).Scan(&student.UUID, &student.Name, &student.Password, &student.Type, &student.RentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &student, nil
}

func (s *service) GetStudentFromName(name string) (*modals.User, error) {
	var student modals.User

	query := "SELECT * FROM users WHERE name = $1 AND type = 0"
	err := s.db.QueryRow(query, name).Scan(&student.UUID, &student.Name, &student.Password, &student.Type, &student.RentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &student, nil
}

func (s *service) DeleteStudentsFromUUID(uuid string) error {
	q := "DELETE FROM users WHERE uuid = $1 AND type = 0"

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

func (s *service) NumberOfRentedBooks(studentUUID string) (int, error) {
	var rentedBooks int

	query := "SELECT rentedbooks FROM users WHERE uuid = $1 AND type = 0"
	err := s.db.QueryRow(query, studentUUID).Scan(&rentedBooks)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrItemNotFound
		} else {
			return 0, err
		}
	}

	return rentedBooks, nil
}

func (s *service) IncreaseStudentRented(studentUUID string) error {
	q := `
    UPDATE users  
    SET rentedbooks = rentedbooks + 1
    WHERE uuid = $1 AND type = 0
  `

	res, err := s.db.Exec(q, studentUUID)
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

func (s *service) DecreaseStudentRented(studentUUID string) error {
	q := `
    UPDATE users  
    SET rentedbooks = rentedbooks - 1
    WHERE uuid = $1
  	AND type = 0
  `

	res, err := s.db.Exec(q, studentUUID)
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
