package database

import (
	"errors"
	"github.com/lib/pq"
)

var ErrItemAlreadyExists = errors.New("Item Already Exists")
var ErrItemNotFound = errors.New("Item Not Found")
var ErrItemMismatch = errors.New("Item Mismatch")
var ErrInternalDatabaseError = errors.New("Internal Database Error")
var ErrInvalidUserType = errors.New("Invalid User Type")

func IsUniqueViolation(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505"
	}
	return false
}

func PrimaryKeyError(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23503"
	}
	return false
}
