package errors

import "errors"

var (
	ErrNewsletterAlreadyExists = errors.New("newsletter already exists")
	ErrNewsletterDoesntExists  = errors.New("newsletter does not exist")
)
