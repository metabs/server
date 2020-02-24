package workspace

import (
	"errors"
	"testing"
)

func TestNewCustomerID(t *testing.T) {

	table := []struct {
		name        string
		raw         string
		wantCustomerID string
		wantErr     error
	}{
		{
			name:        "valid customer id",
			raw:         "c43a5446-b864-4c63-b360-c035ba26057b",
			wantCustomerID: "c43a5446-b864-4c63-b360-c035ba26057b",
			wantErr:     nil,
		},
		{
			name:        "invalid customer id",
			raw:         "z43a5446-x123-4c63-b360-c035ba26057b",
			wantCustomerID: "",
			wantErr:     ErrInvalidCustomerID,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			d, err := NewCustomerID(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if string(d) != r.wantCustomerID {
				t.Error("could not match customer id")
				t.Errorf("want: %s", r.wantCustomerID)
				t.Errorf("got : %s", d)
			}
		})
	}
}
