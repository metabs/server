package customer

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidEmail is used when an invalid email is used
	ErrInvalidEmail = errors.New("customer: could not validate email")

	emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s\.]+\.[^@\.\s]+$`)
)

// Email represents an email
type Email string

// NewEmail return an email and an error back
func NewEmail(e string) (Email, error) {

	if !emailRegex.MatchString(e) {
		return "", ErrInvalidEmail
	}

	return Email(e), nil
}
