package collection

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection/tab"
	"testing"
)

func TestNewCollection(t *testing.T) {
	id := ID("a")
	name := Name("a")
	collection := New(id, name)

	if collection.ID != id {
		t.Error("could not match id")
		t.Errorf("want: %s", id)
		t.Errorf("got : %s", collection.ID)
	}

	if collection.Name != name {
		t.Error("could not match name")
		t.Errorf("want: %s", name)
		t.Errorf("got : %s", collection.Name)
	}

	if collection.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", collection.Created)
	}

	if !collection.Updated.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", collection.Updated)
	}

	name = "b"
	collection.Rename(name)
	if collection.Name != name {
		t.Error("could not match name")
		t.Errorf("want: %s", name)
		t.Errorf("got : %s", collection.Name)
	}
	if collection.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", collection.Created)
	}

	update1 := collection.Updated
	if update1.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", update1)
	}

	tab1 := &tab.Tab{ID: "1"}
	tab2 := &tab.Tab{ID: "2"}
	tabs := []*tab.Tab{tab1, tab2}
	collection.AddTabs(tab1, tab2)

	for i, tb := range collection.Tabs {
		if tb.ID != tabs[i].ID {
			t.Errorf("could not match tab on index %d", i)
			t.Errorf("want: %v", tabs[i])
			t.Errorf("got : %v", tb)
		}
	}

	if len(collection.Tabs) != len(tabs) {
		t.Error("could not match tabs number")
		t.Errorf("want: %d", len(tabs))
		t.Errorf("got : %d", len(collection.Tabs))
	}
	update2 := collection.Updated
	if update2.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", update2)
	}

	if update1.Equal(update2) {
		t.Error("could not mismatch updated time")
		t.Errorf("upadte 1 : %s", update1)
		t.Errorf("upadte 2 : %s", update2)
	}
}
