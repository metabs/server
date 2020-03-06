package workspace

import (
	"context"
	"errors"
	"github.com/metabs/server/workspace/collection"
	"github.com/metabs/server/workspace/collection/tab"
)

var (
	//Errors returned by the repository
	ErrRepoNextID           = errors.New("workspace: could not retrieve next workspace id")
	ErrRepoNextCollectionID = errors.New("workspace: could not retrieve collection next id")
	ErrRepoNextTabID        = errors.New("workspace: could not retrieve next tab id")
	ErrRepoList             = errors.New("workspace: could not list the workspaces")
	ErrNotFound             = errors.New("workspace: could not find the workspace")
	ErrRepoGet              = errors.New("workspace: could not get the workspace")
	ErrRepoAdd              = errors.New("workspace: could not add the workspace")
	ErrRepoDelete           = errors.New("workspace: could not delete the workspace")
)

// Repo represents the persistence layer for the workspace aggregate
type Repo interface {
	NextID(context.Context) (ID, error)
	NextCollectionID(context.Context) (collection.ID, error)
	NextTabID(context.Context) (tab.ID, error)
	List(context.Context, CustomerID) ([]*Workspace, error)
	Get(context.Context, ID) (*Workspace, error)
	Add(context.Context, *Workspace) error
	Delete(context.Context, ID) error
}
