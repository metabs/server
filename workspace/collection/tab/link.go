package tab

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidLink is used when an invalid link is given
var ErrInvalidLink = errors.New("tab: could not use invalid link")

// Link represents a tab link
type Link string

// NewLink returns a link and an error back
func NewLink(i string) (Link, error) {
	_, err := url.ParseRequestURI(i)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidLink, err)
	}

	return Link(i), nil
}

// String returns a string representation of the link
func (l Link) String() string {
	return string(l)
}