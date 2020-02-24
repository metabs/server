package workspace

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"testing"
)

func TestNewWorkspace(t *testing.T) {
	id := ID("a")
	name := Name("a")
	customerID := CustomerID("a")
	workspace := New(id, name, customerID)

	if workspace.ID != id {
		t.Error("could not match id")
		t.Errorf("want: %s", id)
		t.Errorf("got : %s", workspace.ID)
	}

	if workspace.Name != name {
		t.Error("could not match name")
		t.Errorf("want: %s", name)
		t.Errorf("got : %s", workspace.Name)
	}

	if workspace.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", workspace.Created)
	}

	if !workspace.Updated.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", workspace.Updated)
	}

	name = "b"
	workspace.Rename(name)
	if workspace.Name != name {
		t.Error("could not match name")
		t.Errorf("want: %s", name)
		t.Errorf("got : %s", workspace.Name)
	}
	if workspace.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", workspace.Created)
	}

	update1 := workspace.Updated
	if update1.IsZero() {
		t.Error("could not find a updated time")
		t.Errorf("got : %s", update1)
	}

	collection1 := collection.Collection{ID: "1"}
	collection2 := collection.Collection{ID: "2"}
	collections := []collection.Collection{collection1, collection2}
	workspace.AddCollections(collection1, collection2)

	for i, c := range workspace.Collections {
		if c.ID != collections[i].ID {
			t.Errorf("could not match collection on index %d", i)
			t.Errorf("want: %v", collections[i])
			t.Errorf("got : %v", c)
		}
	}

	if len(workspace.Collections) != len(collections) {
		t.Error("could not match collections number")
		t.Errorf("want: %d", len(collections))
		t.Errorf("got : %d", len(workspace.Collections))
	}
	update2 := workspace.Updated
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
