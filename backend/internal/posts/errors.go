package posts

import "errors"

var (
	ErrPostAlreadyExists = errors.New("post already exists")
	ErrPostNotFound      = errors.New("post not found")
	InvalidRequstBody    = errors.New("title and/or description must be filled")
)
