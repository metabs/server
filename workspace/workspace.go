package workspace

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"time"
)

// Workspace represent a workspace
type Workspace struct {
	ID          ID                      `json:"id,string"`
	Name        Name                    `json:"name,string"`
	OwnerID     OwnerID                 `json:"owner_id,string"`
	Collections []collection.Collection `json:"collections"`
	Created     time.Time               `json:"created,string"`
	Updated     time.Time               `json:"updated,string"`
}

// New returns a new workspace created for the first time
func New(
	id ID,
	name Name,
	ownerID OwnerID,
) *Workspace {
	return &Workspace{
		ID:      id,
		Name:    name,
		OwnerID: ownerID,
		Created: time.Now(),
	}
}

// Rename change the name of a workspace
func (w *Workspace) Rename(name Name) {
	w.Name = name
	w.Updated = time.Now()
}

// AddCollections change the name of a workspace
func (w *Workspace) AddCollections(collections ...collection.Collection) {
	w.Collections = append(w.Collections, collections...)
	w.Updated = time.Now()
}
