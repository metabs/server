package workspace

import (
	"context"
	"errors"
)

var (
	//Errors that can be used from the repository
	ErrRepoNextID = errors.New("workspace: could not retrieve next id for the workspace")
	ErrRepoList   = errors.New("workspace: could not list the workspaces")
	ErrNotFound   = errors.New("workspace: could not find the workspace")
	ErrRepoGet    = errors.New("workspace: could not get the workspace")
	ErrRepoAdd    = errors.New("workspace: could not add the workspace")
	ErrRepoDelete = errors.New("workspace: could not delete the workspace")
)

// Repo represents the persistence layer for the workspace aggregate
type Repo interface {
	NextID(context.Context) (ID, error)
	List(context.Context, CustomerID) ([]*Workspace, error)
	Get(context.Context, ID) (*Workspace, error)
	Add(context.Context, *Workspace) error
	Delete(context.Context, ID) error
}
