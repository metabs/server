package collection

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minNameLength = 1
	maxNameLength = 50
)

var (
	// Errors used when an invalid name is given
	ErrInvalidName  = errors.New("collection: could not use invalid name")
	ErrNameTooShort = fmt.Errorf("%w: min length allowed is %d", ErrInvalidName, minNameLength)
	ErrNameTooLong  = fmt.Errorf("%w: max length allowed is %d", ErrInvalidName, maxNameLength)
)

// Name represents a collection name
type Name string

// NewName returns an name and an error back
func NewName(d string) (Name, error) {
	switch l := len(strings.TrimSpace(d)); {
	case l < minNameLength:
		return "", ErrNameTooShort
	case l > maxNameLength:
		return "", ErrNameTooLong
	default:
		return Name(d), nil
	}
}

// String returns a string representation of the name
func (n Name) String() string {
	return string(n)
}
