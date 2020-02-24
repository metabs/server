package workspace

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// ErrInvalidOwnerID is used when an invalid id is used
var ErrInvalidOwnerID = errors.New("workspace: could not validate id")

// OwnerID represents a workspace owner id
type OwnerID string

// NewOwnerID return an id and an error back
func NewOwnerID(id string) (OwnerID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", fmt.Errorf("%w: '%s' is not valide due to: %s", ErrInvalidOwnerID, id, err)
	}

	return OwnerID(id), nil
}
