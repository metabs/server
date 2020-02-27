package customer

import (
	"context"
	"errors"
)

var (
	//Errors returned by the repository
	ErrRepoNextID = errors.New("customer: could not retrieve next customer id")
	ErrNotFound   = errors.New("customer: could not find the customer")
	ErrRepoGet    = errors.New("customer: could not get the customer")
	ErrRepoAdd    = errors.New("customer: could not add the customer")
	ErrRepoDelete = errors.New("customer: could not delete the customer")
)

// Repo represents the persistence layer for the customer aggregate
type Repo interface {
	NextID(context.Context) (ID, error)
	Get(context.Context, ID) (*Customer, error)
	GetByEmail(context.Context, Email) (*Customer, error)
	Add(context.Context, *Customer) error
	Delete(context.Context, ID) error
}
