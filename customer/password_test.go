package customer

import (
	"errors"
	"testing"
)

func TestNewPassword(t *testing.T) {

	table := []struct {
		name      string
		raw       string
		matchRaw  string
		wantMatch bool
		wantErr   error
	}{
		{
			name:      "valid password matching",
			raw:       "&dsh_10wMdisjl^0293",
			matchRaw:  "&dsh_10wMdisjl^0293",
			wantMatch: true,
			wantErr:   nil,
		},
		{
			name:      "valid password not matching",
			raw:       "&dsh_10wMdisjl^0293",
			matchRaw:  "&nopedsh_10wMdisjl^0293",
			wantMatch: false,
			wantErr:   nil,
		},
		{
			name:      "password too simple",
			raw:       "password",
			matchRaw:  "password",
			wantMatch: false,
			wantErr:   ErrPasswordTooSimple,
		},
		{
			name:      "password too long",
			raw:       "passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassw",
			matchRaw:  "password",
			wantMatch: false,
			wantErr:   ErrPasswordTooLong,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			pwd, err := NewPassword(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if pwd.Compare(r.matchRaw) != r.wantMatch {
				t.Error("could not match password comparision")
				t.Errorf("want: %t", r.wantMatch)
			}
		})
	}
}
