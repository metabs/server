package tab

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidLink is used when an invalid link is used
var ErrInvalidLink = errors.New("tab: could not validate link")

// Link represents a tab link
type Link url.URL

// NewLink return an link and an error back
func NewLink(i string) (Link, error) {
	u, err := url.ParseRequestURI(i)
	if err != nil {
		return Link{}, fmt.Errorf("%w: '%s' is not valide due to: %s", ErrInvalidLink, i, err)
	}

	return Link(*u), nil
}
