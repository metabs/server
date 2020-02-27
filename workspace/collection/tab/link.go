package tab

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidLink is used when an invalid link is used
var ErrInvalidLink = errors.New("tab: could not validate link")

// Link represents a tab link
type Link string

// NewLink return an link and an error back
func NewLink(i string) (Link, error) {
	_, err := url.ParseRequestURI(i)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidLink, err)
	}

	return Link(i), nil
}
