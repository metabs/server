package workspace

import (
	"errors"
	"testing"
)

func TestNewOwnerID(t *testing.T) {

	table := []struct {
		name        string
		raw         string
		wantOwnerID string
		wantErr     error
	}{
		{
			name:        "valid owner id",
			raw:         "c43a5446-b864-4c63-b360-c035ba26057b",
			wantOwnerID: "c43a5446-b864-4c63-b360-c035ba26057b",
			wantErr:     nil,
		},
		{
			name:        "invalid owner id",
			raw:         "z43a5446-x123-4c63-b360-c035ba26057b",
			wantOwnerID: "",
			wantErr:     ErrInvalidOwnerID,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			d, err := NewOwnerID(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if string(d) != r.wantOwnerID {
				t.Error("could not match owner id")
				t.Errorf("want: %s", r.wantOwnerID)
				t.Errorf("got : %s", d)
			}
		})
	}
}
