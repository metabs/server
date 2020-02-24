package workspace

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minNameLength = 1
	maxNameLength = 75
)

// ErrInvalidName is used when an invalid name is used
var (
	ErrInvalidName  = errors.New("workspace: could not validate name")
	ErrNameTooShort = fmt.Errorf("%w: min length allowed is %d", ErrInvalidName, minNameLength)
	ErrNameTooLong  = fmt.Errorf("%w: max length allowed is %d", ErrInvalidName, maxNameLength)
)

// Name represents a workspace name
type Name string

// NewName return an name and an error back
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
