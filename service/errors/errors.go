package errors

import "errors"

var (
	ErrForbidden = errors.New("You do not have permission to perform this action")
	ErrUnauthorized = errors.New("You are not authorized to perform this action")
	ErrInternalServerError = errors.New("Internal server error")
	ErrBadRequest = errors.New("Bad request")
	ErrNotFound = errors.New("Not found")
)
