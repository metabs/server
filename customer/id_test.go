package customer

import (
	"errors"
	"testing"
)

func TestNewID(t *testing.T) {

	table := []struct {
		name    string
		raw     string
		wantID  string
		wantErr error
	}{
		{
			name:    "valid customer id",
			raw:     "c43a5446-b864-4c63-b360-c035ba26057b",
			wantID:  "c43a5446-b864-4c63-b360-c035ba26057b",
			wantErr: nil,
		},
		{
			name:    "invalid customer id",
			raw:     "z43a5446-x123-4c63-b360-c035ba26057b",
			wantID:  "",
			wantErr: ErrInvalidID,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			d, err := NewID(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if string(d) != r.wantID {
				t.Error("could not match customer id")
				t.Errorf("want: %s", r.wantID)
				t.Errorf("got : %s", d)
			}
		})
	}
}
