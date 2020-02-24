package collection

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection/tab"
	"time"
)

// Collection represent a collection
type Collection struct {
	ID      ID        `json:"id,string"`
	Name    Name      `json:"name,string"`
	Tabs    []tab.Tab `json:"tabs"`
	Created time.Time `json:"created,string"`
	Updated time.Time `json:"updated,string"`
}

// New returns a new collection created for the first time
func New(id ID, name Name) *Collection {
	return &Collection{
		ID:      id,
		Name:    name,
		Tabs:    make([]tab.Tab, 0),
		Created: time.Now(),
	}
}

// Rename renames a collection
func (c *Collection) Rename(name Name) {
	c.Name = name
	c.Updated = time.Now()
}

// AddTabs adds tabs to the collection
func (c *Collection) AddTabs(tab ...tab.Tab) {
	c.Tabs = append(c.Tabs, tab...)
	c.Updated = time.Now()
}
