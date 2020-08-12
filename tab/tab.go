package tab

import (
	"time"
)

// Tab represents a tab
type Tab struct {
	ID          ID          `json:"id"`
	Title       Title       `json:"title"`
	Description Description `json:"description"`
	Icon        Icon        `json:"icon"`
	Link        Link        `json:"link"`
	Created     time.Time   `json:"created"`
	Updated     time.Time   `json:"updated,omitempty"`
}

// New returns a tab created for the first time
func New(id ID, title Title, description Description, icon Icon, link Link) *Tab {
	return &Tab{
		ID:          id,
		Title:       title,
		Description: description,
		Icon:        icon,
		Link:        link,
		Created:     time.Now(),
	}
}

// Update updates a tab with new identifiers
func (t *Tab) Update(title Title, description Description, icon Icon, link Link) {
	t.Title = title
	t.Description = description
	t.Icon = icon
	t.Link = link
	t.Updated = time.Now()
}
