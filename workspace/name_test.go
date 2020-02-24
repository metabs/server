package workspace

import (
	"errors"
	"testing"
)

func TestNewName(t *testing.T) {

	table := []struct {
		name     string
		raw      string
		wantName string
		wantErr  error
	}{
		{
			name:     "valid name",
			raw:      "this is a name for my tab",
			wantName: "this is a name for my tab",
			wantErr:  nil,
		},
		{
			name:     "too long name",
			raw:      "this is a too long name, this is a too long name, this is a too long name, this is a too long name, this is a too long name!",
			wantName: "",
			wantErr:  ErrNameTooLong,
		},
		{
			name:     "empty name",
			raw:      "",
			wantName: "",
			wantErr:  ErrNameTooShort,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			d, err := NewName(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if string(d) != r.wantName {
				t.Error("could not match name")
				t.Errorf("want: %s", r.wantName)
				t.Errorf("got : %s", d)
			}
		})
	}
}
