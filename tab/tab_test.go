package tab

import (
	"testing"
)

func TestNewTab(t *testing.T) {
	id := ID("a")
	title := Title("a")
	description := Description("a")
	icon := Icon("https://github.com/damianopetrungaro/profile.png")
	link := Link("http://github.com/damianopetrungaro/profile.png")

	tab := New(id, title, description, icon, link)
	if tab.ID != id {
		t.Error("could not match id")
		t.Errorf("want: %s", id)
		t.Errorf("got : %s", tab.ID)
	}

	if tab.ID != id {
		t.Error("could not match id")
		t.Errorf("want: %s", id)
		t.Errorf("got : %s", tab.ID)
	}

	if tab.Title != title {
		t.Error("could not match title")
		t.Errorf("want: %s", title)
		t.Errorf("got : %s", tab.Title)
	}

	if tab.Description != description {
		t.Error("could not match description")
		t.Errorf("want: %s", description)
		t.Errorf("got : %s", tab.Description)
	}

	if tab.Icon != icon {
		t.Error("could not match icon")
		t.Errorf("want: %v", icon)
		t.Errorf("got : %v", tab.Icon)
	}

	if tab.Link != link {
		t.Error("could not match link")
		t.Errorf("want: %v", link)
		t.Errorf("got : %v", tab.Link)
	}

	if tab.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", tab.Created)
	}

	if !tab.Updated.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", tab.Updated)
	}

	title = "b"
	description = "b"
	icon = "https://www.github.com/damianopetrungaro/image.png"
	link = "http://www.github.com/damianopetrungaro/image.png"

	tab.Update(title, description, icon, link)

	if tab.ID != id {
		t.Error("could not match id")
		t.Errorf("want: %s", id)
		t.Errorf("got : %s", tab.ID)
	}

	if tab.Title != title {
		t.Error("could not match title")
		t.Errorf("want: %s", title)
		t.Errorf("got : %s", tab.Title)
	}

	if tab.Description != description {
		t.Error("could not match description")
		t.Errorf("want: %s", description)
		t.Errorf("got : %s", tab.Description)
	}

	if tab.Icon != icon {
		t.Error("could not match icon")
		t.Errorf("want: %v", icon)
		t.Errorf("got : %v", tab.Icon)
	}

	if tab.Link != link {
		t.Error("could not match link")
		t.Errorf("want: %v", link)
		t.Errorf("got : %v", tab.Link)
	}

	if tab.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", tab.Created)
	}

	if tab.Updated.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", tab.Updated)
	}
}
