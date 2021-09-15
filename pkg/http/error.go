package http

import "errors"

var (
	errUserInternal   = errors.New("an internal error has occurred")
	errInvalidQueries = errors.New("failed to validate input queries")
	errStorage        = errors.New("failed to perform storage operation")
)
