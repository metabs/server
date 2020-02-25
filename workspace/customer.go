package workspace

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// ErrInvalidCustomerID is used when an invalid id is used
var ErrInvalidCustomerID = errors.New("workspace: could not validate id")

// CustomerID represents a workspace customer id
type CustomerID string

// NewCustomerID return an id and an error back
func NewCustomerID(id string) (CustomerID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", fmt.Errorf("%w: '%s' is not valide due to: %s", ErrInvalidCustomerID, id, err)
	}

	return CustomerID(id), nil
}


func (c *CustomerID) UnmarshalJSON(data []byte) error {
	c2, err := NewCustomerID(string(data))
	if err != nil {
		return err
	}
	*c = c2

	return nil
}

func (c CustomerID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c)), nil
}
