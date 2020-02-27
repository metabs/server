package tab

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minTitleLength = 1
	maxTitleLength = 50
)

// ErrInvalidTitle is used when an invalid title is used
var (
	ErrInvalidTitle  = errors.New("tab: could not validate title")
	ErrTitleTooShort = fmt.Errorf("%w: min length allowed is %d", ErrInvalidTitle, minTitleLength)
	ErrTitleTooLong  = fmt.Errorf("%w: max length allowed is %d", ErrInvalidTitle, maxTitleLength)
)

// Title represents a tab title
type Title string

// NewTitle return an title and an error back
func NewTitle(d string) (Title, error) {
	switch l := len(strings.TrimSpace(d)); {
	case l < minTitleLength:
		return "", ErrTitleTooShort
	case l > maxTitleLength:
		return "", ErrTitleTooLong
	default:
		return Title(d), nil
	}
}