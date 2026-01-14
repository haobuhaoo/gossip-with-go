package posts

import "errors"

var (
	ErrPostAlreadyExists = errors.New("post already exists")
	ErrPostNotFound      = errors.New("post not found")
)
