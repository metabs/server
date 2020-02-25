package workspace

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"time"
)

// Workspace represent a workspace
type Workspace struct {
	ID          ID                       `json:"id,string"`
	Name        Name                     `json:"name,string"`
	CustomerID  CustomerID               `json:"customer_id,string"`
	Collections []*collection.Collection `json:"collections"`
	Created     time.Time                `json:"created"`
	Updated     time.Time                `json:"updated"`
}

// New returns a new workspace created for the first time
func New(
	id ID,
	name Name,
	customerID CustomerID,
) *Workspace {
	return &Workspace{
		ID:          id,
		Name:        name,
		CustomerID:  customerID,
		Collections: make([]*collection.Collection, 0),
		Created:     time.Now(),
	}
}

// Rename change the name of a workspace
func (w *Workspace) Rename(name Name) {
	w.Name = name
	w.Updated = time.Now()
}

// AddCollections change the name of a workspace
func (w *Workspace) AddCollections(collections ...*collection.Collection) {
	w.Collections = append(w.Collections, collections...)
	w.Updated = time.Now()
}
