package collection

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection/tab"
	"time"
)

// Collection represent a collection
type Collection struct {
	ID      ID         `json:"id"`
	Name    Name       `json:"name"`
	Tabs    []*tab.Tab `json:"tabs,omitempty"`
	Created time.Time  `json:"created"`
	Updated time.Time  `json:"updated,omitempty"`
}

// New returns a collection created for the first time
func New(id ID, name Name) *Collection {
	return &Collection{
		ID:      id,
		Name:    name,
		Tabs:    make([]*tab.Tab, 0),
		Created: time.Now(),
	}
}

// Rename renames a collection
func (c *Collection) Rename(name Name) {
	c.Name = name
	c.Updated = time.Now()
}

// AddTabs adds tabs to the collection
func (c *Collection) AddTabs(tabs ...*tab.Tab) {
	c.Tabs = append(c.Tabs, tabs...)
	c.Updated = time.Now()
}

// RemoveTab removes a tab if it exists
func (c *Collection) RemoveTab(id tab.ID) bool {
	for i, t := range c.Tabs {
		if t.ID == id {
			c.Tabs[i] = c.Tabs[len(c.Tabs)-1]
			c.Tabs[len(c.Tabs)-1] = nil
			c.Tabs = c.Tabs[:len(c.Tabs)-1]
			c.Updated = time.Now()
			return true
		}
	}

	return false
}

// FindTab returns a tab if it exists
func (c *Collection) FindTab(id tab.ID) (*tab.Tab, bool) {
	for _, t := range c.Tabs {
		if t.ID == id {
			return t, true
		}
	}

	return nil, false
}

// UpdateTab updates a tab if it exists
func (c *Collection) UpdateTab(t *tab.Tab) bool {
	for i, tb := range c.Tabs {
		if tb.ID == t.ID {
			c.Tabs[i] = t
			c.Updated = time.Now()
			return true
		}
	}

	return false
}
