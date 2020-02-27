package customer

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 10
	maxPasswordLength = 100
)

var (
	// Errors used when an invalid password is given
	ErrInvalidPassword   = errors.New("customer: could not use invalid password")
	ErrPasswordTooLong   = fmt.Errorf("%w: too long", ErrInvalidPassword)
	ErrPasswordTooSimple = fmt.Errorf("%w: too simple", ErrInvalidPassword)
)

// Password represents an password
type Password string

// NewPassword returns a password and an error back
func NewPassword(pwd string) (Password, error) {
	if len(pwd) < minPasswordLength {
		return "", ErrPasswordTooSimple
	}

	if len(pwd) > maxPasswordLength {
		return "", ErrPasswordTooLong
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidPassword, err)
	}

	return Password(string(hash)), nil
}

// Compare return true if passwords are matching
func (p Password) Compare(psw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(psw)) == nil
}

// String returns a string representation of the password
func (p Password) String() string {
	return string(p)
}
