package auth

import "errors"

var ErrWrongPasssword = errors.New("Wrong Password")
var ErrInvalidUserName = errors.New("Invalid Username")
