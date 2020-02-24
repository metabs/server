package tab

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// ErrInvalidID is used when an invalid id is used
var ErrInvalidID = errors.New("tab: could not validate id")

// ID represents a tab id
type ID string

// NewID return an id and an error back
func NewID(id string) (ID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidID, err)
	}

	return ID(id), nil
}
