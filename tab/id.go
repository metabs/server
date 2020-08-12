package tab

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// ErrInvalidID is used when an invalid id is given
var ErrInvalidID = errors.New("tab: could not use invalid id")

// ID represents a tab id
type ID string

// NewID returns an id and an error back
func NewID(id string) (ID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidID, err)
	}

	return ID(id), nil
}

// String returns a string representation of the id
func (i ID) String() string {
	return string(i)
}