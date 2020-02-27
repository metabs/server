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

// ErrInvalidDescription is used when an invalid description is used
var (
	ErrInvalidDescription  = errors.New("tab: could not validate description")
	ErrDescriptionTooShort = fmt.Errorf("%w: min length allowed is %d", ErrInvalidDescription, minDescriptionLength)
	ErrDescriptionTooLong  = fmt.Errorf("%w: max length allowed is %d", ErrInvalidDescription, maxDescriptionLength)
)

// Description represents a tab description
type Description string

// NewDescription return an description and an error back
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