package collection

import "context"

// Repo represents the persistence layer for the collection aggregate
type Repo interface {
	Get(context.Context, ID) (*Collection, error)
	Add(context.Context, *Collection) error
	Delete(context.Context, ID) error
}
