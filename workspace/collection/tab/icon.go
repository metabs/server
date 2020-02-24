package tab

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidIcon is used when an invalid icon is used
var ErrInvalidIcon = errors.New("tab: could not validate icon")

// Icon represents a tab icon
type Icon url.URL

// NewIcon return an icon and an error back
func NewIcon(i string) (Icon, error) {
	u, err := url.ParseRequestURI(i)
	if err != nil {
		return Icon{}, fmt.Errorf("%w: '%s", ErrInvalidIcon, err)
	}

	return Icon(*u), nil
}
