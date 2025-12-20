package users

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exist")
	ErrUserNotFound      = errors.New("user not found")
)
