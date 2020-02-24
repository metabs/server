package workspace

import "context"

// Repo represents the persistence layer for the workspace aggregate
type Repo interface {
	List(context.Context, CustomerID) ([]*Workspace, error)
	Get(context.Context, ID) (*Workspace, error)
	Add(context.Context, *Workspace) error
	Delete(context.Context, ID) error
}
