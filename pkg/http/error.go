package http

import "errors"

var (
	errUserInternal = errors.New("an internal error has occurred")
	errInvalidQuery = errors.New("failed to validate input query string")
	errStorage      = errors.New("failed to perform storage operation")
)
