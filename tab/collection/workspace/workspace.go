package workspace

import (
	"github.com/metabs/server/tab"
	"github.com/metabs/server/tab/collection"
	"time"
)

// Workspace represent a workspace
type Workspace struct {
	ID          ID                       `json:"id"`
	Name        Name                     `json:"name"`
	CustomerID  CustomerID               `json:"customer_id"`
	Collections []*collection.Collection `json:"collections,omitempty"`
	Created     time.Time                `json:"created"`
	Updated     time.Time                `json:"updated,omitempty"`
}

// New returns a workspace created for the first time
func New(id ID, name Name, customerID CustomerID) *Workspace {
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

// AddCollections add a collection
func (w *Workspace) AddCollections(collections ...*collection.Collection) {
	w.Collections = append(w.Collections, collections...)
	w.Updated = time.Now()
}

// RemoveCollection removes a collection if it exists
func (w *Workspace) RemoveCollection(id collection.ID) bool {
	for i, coll := range w.Collections {
		if coll.ID == id {
			w.Collections[i] = w.Collections[len(w.Collections)-1]
			w.Collections[len(w.Collections)-1] = nil
			w.Collections = w.Collections[:len(w.Collections)-1]
			w.Updated = time.Now()
			return true
		}
	}

	return false
}

// RenameCollection renames a collection if it exists
func (w *Workspace) RenameCollection(id collection.ID, name collection.Name) bool {
	for _, coll := range w.Collections {
		if coll.ID == id {
			coll.Rename(name)
			w.Updated = time.Now()
			return true
		}
	}

	return false
}

// AddTabs adds tabs to a collection
func (w *Workspace) AddTabs(id collection.ID, tabs ...*tab.Tab) bool {
	for _, coll := range w.Collections {
		if coll.ID == id {
			coll.AddTabs(tabs...)
			w.Updated = time.Now()
			return true
		}
	}

	return false
}

// RemoveTab removes a tab from a collection if it exists
func (w *Workspace) RemoveTab(id tab.ID, collID collection.ID) bool {
	for _, coll := range w.Collections {
		if coll.ID == collID {
			if !coll.RemoveTab(id) {
				return false
			}
			w.Updated = time.Now()
			return true
		}
	}

	return false
}

// UpdateTab updates a tab of a collection
func (w *Workspace) UpdateTab(t *tab.Tab, collID collection.ID) bool {
	for _, coll := range w.Collections {
		if coll.ID == collID {
			if !coll.UpdateTab(t) {
				return false
			}
			w.Updated = time.Now()
			return true
		}
	}

	return false
}

// FindTab returns a tab from a collection if it exists
func (w *Workspace) FindTab(id tab.ID, collID collection.ID) (*tab.Tab, bool) {
	for _, coll := range w.Collections {
		if coll.ID == collID {
			return coll.FindTab(id)
		}
	}

	return nil, false
}
