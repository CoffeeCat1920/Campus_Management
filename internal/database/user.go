package database

import modals "what/internal/models"

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
