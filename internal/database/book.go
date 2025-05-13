package database

import (
	"database/sql"
	modals "what/internal/models"
)

func (s *service) AddBook(book *modals.Book) error {
	q := `
  INSERT INTO books(uuid, isbn, name, available)
  VALUES($1, $2, $3, $4)
  `
	_, err := s.db.Exec(q, book.UUID, book.ISBN, book.Name, book.Available)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateBookName(uuid string, name string) error {
	q := `
	UPDATE books
	SET name = $1
	WHERE uuid = $2
  `
	result, err := s.db.Exec(q, name, uuid)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrItemNotFound
	}

	return nil
}

func (s *service) DeleteBook(uuid string) error {

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

func (s *service) IsAvailable(uuid string) (bool, error) {
	var available bool
	q := "SELECT available FROM books where uuid = $1"

	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&available)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, ErrItemNotFound
		} else {
			return false, err
		}
	}

	return available, nil
}

func (s *service) ToggleBookAvailiablity(uuid string) error {
	q := `
	UPDATE books
	SET available = NOT available
	WHERE uuid = $1;
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

func (s *service) SearchBooks(name string) ([]modals.Book, error) {
	var recipes []modals.Book

	searchTerm := "%" + name + "%"
	query := "SELECT * FROM books WHERE name ILIKE $1"

	rows, err := s.db.Query(query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe modals.Book
		err := rows.Scan(&recipe.UUID, &recipe.ISBN, &recipe.Name, &recipe.Available)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (s *service) GetBookUUIDFromISBN(isbn string) (string, error) {
	var uuid string

	q := `SELECT uuid FROM books 
	WHERE isbn = $1`

	row := s.db.QueryRow(q, isbn)
	err := row.Scan(&uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrItemNotFound
		} else {
			return "", err
		}
	}

	return uuid, err
}

func (s *service) GetBookFromUUID(uuid string) (*modals.Book, error) {
	var book modals.Book

	q := `SELECT * FROM books 
	WHERE uuid = $1`

	row := s.db.QueryRow(q, uuid)
	err := row.Scan(&book.UUID, &book.ISBN, &book.Name, &book.Available)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		} else {
			return nil, err
		}
	}

	return &book, err
}

func (s *service) GetAllBooks() ([]modals.Book, error) {
	recipes := []modals.Book{}

	query := "SELECT * FROM books ORDER BY name ASC"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe modals.Book
		err := rows.Scan(&recipe.UUID, &recipe.ISBN, &recipe.Name, &recipe.Available)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (s *service) GetAllBooksExcept() ([]modals.Book, error) {
	recipes := []modals.Book{}
	return recipes, nil
}
