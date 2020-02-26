package collection

import (
	"context"
	"errors"
)

var (
	//Errors that can be used from the repository
	ErrRepoNextID = errors.New("collection: could not retrieve next id for the collection")
)

// Repo represents the persistence layer for the collection aggregate
type Repo interface {
	NextID(context.Context) (ID, error)
}
