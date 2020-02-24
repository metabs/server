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
	// ErrInvalidPassword is used when an invalid password is used
	ErrInvalidPassword   = errors.New("customer: could not validate password")
	ErrPasswordTooLong   = fmt.Errorf("%w: too long", ErrInvalidPassword)
	ErrPasswordTooSimple = fmt.Errorf("%w: too simple", ErrInvalidPassword)
)

// Password represents an password
type Password string

// NewPassword return an password and an error back
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
