package tab

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidIcon is used when an invalid icon is given
var ErrInvalidIcon = errors.New("tab: could not use invalid icon")

// Icon represents a tab icon
type Icon string

// NewIcon returns an icon and an error back
func NewIcon(i string) (Icon, error) {
	_, err := url.ParseRequestURI(i)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidIcon, err)
	}

	return Icon(i), nil
}

// String returns a string representation of the icon
func (i Icon) String() string {
	return string(i)
}