package customer

import (
	"errors"
	"testing"
)

func TestNewEmail(t *testing.T) {

	table := []struct {
		name      string
		raw       string
		wantEmail string
		wantErr   error
	}{
		{
			name:      "valid email",
			raw:       "thedam.petr@hotmail.com",
			wantEmail: "thedam.petr@hotmail.com",
			wantErr:   nil,
		},
		{
			name:      "invalid email",
			raw:       "per.com",
			wantEmail: "",
			wantErr:   ErrInvalidEmail,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			e, err := NewEmail(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}

			if e.String() != r.wantEmail {
				t.Error("could not match email")
				t.Errorf("want: %s", r.wantEmail)
				t.Errorf("got : %s", e)
			}
		})
	}
}
