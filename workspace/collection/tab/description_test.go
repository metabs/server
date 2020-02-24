package tab

import (
	"errors"
	"testing"
)

func TestNewDescription(t *testing.T) {

	table := []struct {
		name     string
		raw      string
		wantDesc string
		wantErr  error
	}{
		{
			name:     "valid description",
			raw:      "this is a description for my tab",
			wantDesc: "this is a description for my tab",
			wantErr:  nil,
		},
		{
			name:     "too long description",
			raw:      "this is a too long description, this is a too long description, this is a too long description, this is a too long description, this is a too long description, this is a too long description",
			wantDesc: "",
			wantErr:  ErrDescriptionTooLong,
		},
		{
			name:     "empty description",
			raw:      "",
			wantDesc: "",
			wantErr:  ErrDescriptionTooShort,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			d, err := NewDescription(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if string(d) != r.wantDesc {
				t.Error("could not match description")
				t.Errorf("want: %s", r.wantDesc)
				t.Errorf("got : %s", d)
			}
		})
	}
}
