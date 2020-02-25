package workspace

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// ErrInvalidID is used when an invalid id is used
var ErrInvalidID = errors.New("workspace: could not validate id")

// ID represents a workspace id
type ID string

// NewID return an id and an error back
func NewID(id string) (ID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", fmt.Errorf("%w: '%s' is not valide due to: %s", ErrInvalidID, id, err)
	}

	return ID(id), nil
}

func (i *ID) UnmarshalJSON(data []byte) error {
	i2, err := NewID(string(data))
	if err != nil {
		return err
	}
	*i = i2

	return nil
}

func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, i)), nil
}
