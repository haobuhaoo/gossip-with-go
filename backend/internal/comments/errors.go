package comments

import "errors"

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrPostNotUpdated  = errors.New("post not updated")
	ErrVoteNotFound    = errors.New("vote not found")
)
