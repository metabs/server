package tab

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidIcon is used when an invalid icon is used
var ErrInvalidIcon = errors.New("tab: could not validate icon")

// Icon represents a tab icon
type Icon string

// NewIcon return an icon and an error back
func NewIcon(i string) (Icon, error) {
	_, err := url.ParseRequestURI(i)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidIcon, err)
	}

	return Icon(i), nil
}