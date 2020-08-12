package tab

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minDescriptionLength = 1
	maxDescriptionLength = 150
)

var (
	// Errors used when an invalid description is given
	ErrInvalidDescription  = errors.New("tab: could not use invalid description")
	ErrDescriptionTooShort = fmt.Errorf("%w: min length allowed is %d", ErrInvalidDescription, minDescriptionLength)
	ErrDescriptionTooLong  = fmt.Errorf("%w: max length allowed is %d", ErrInvalidDescription, maxDescriptionLength)
)

// Description represents a tab description
type Description string

// NewDescription returns a description and an error back
func NewDescription(d string) (Description, error) {
	switch l := len(strings.TrimSpace(d)); {
	case l < minDescriptionLength:
		return "", ErrDescriptionTooShort
	case l > maxDescriptionLength:
		return "", ErrDescriptionTooLong
	default:
		return Description(d), nil
	}
}

// String returns a string representation of the description
func (d Description) String() string {
	return string(d)
}
