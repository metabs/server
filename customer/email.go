package customer

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidEmail is used when an invalid email is given
	ErrInvalidEmail = errors.New("customer: could not use invalid email")

	emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s\.]+\.[^@\.\s]+$`)
)

// Email represents an email
type Email string

// NewEmail returns an email and an error back
func NewEmail(e string) (Email, error) {

	if !emailRegex.MatchString(e) {
		return "", ErrInvalidEmail
	}

	return Email(e), nil
}

// String returns a string representation of the email
func (e Email) String() string {
	return string(e)
}
