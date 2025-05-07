package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) AddStudent(student *modals.User) error {
	q := `
  INSERT INTO students(uuid, name, password, type, rentedBooks)
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

func (s *service) GetStudentUUIDFromName(name string) (string, error) {
	var studentName string

	query := "SELECT uuid FROM students WHERE name = $1"
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

func (s *service) GetStudentFromName(name string) (*modals.User, error) {
	var student modals.User

	query := "SELECT * FROM students WHERE name = $1"
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

func (s *service) NumberOfRentedBooks(studentUUID string) (int, error) {
	var rentedBooks int

	query := "SELECT rentedbooks FROM students WHERE name = $1"
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
    UPDATE students 
    SET rentedbooks = rentedbooks + 1
    WHERE uuid = $1
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
    UPDATE students 
    SET rentedbooks = rentedbooks - 1
    WHERE uuid = $1
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
