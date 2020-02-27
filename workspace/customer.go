package workspace

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// ErrInvalidCustomerID is used when an invalid customer id is given
var ErrInvalidCustomerID = errors.New("workspace: could not use invalid id")

// CustomerID represents a workspace customer id
type CustomerID string

// NewCustomerID return an id and an error back
func NewCustomerID(id string) (CustomerID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidCustomerID, err)
	}

	return CustomerID(id), nil
}

// String returns a string representation of the customer id
func (i CustomerID) String() string {
	return string(i)
}
