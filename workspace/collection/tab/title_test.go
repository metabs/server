package tab

import (
	"errors"
	"testing"
)

func TestNewTitle(t *testing.T) {

	table := []struct {
		name      string
		raw       string
		wantTitle string
		wantErr   error
	}{
		{
			name:      "valid title",
			raw:       "this is a title for my tab",
			wantTitle: "this is a title for my tab",
			wantErr:   nil,
		},
		{
			name:      "too long title",
			raw:       "this is a too long title, this is a too long title!",
			wantTitle: "",
			wantErr:   ErrTitleTooLong,
		},
		{
			name:      "empty title",
			raw:       "",
			wantTitle: "",
			wantErr:   ErrTitleTooShort,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			d, err := NewTitle(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}
			if string(d) != r.wantTitle {
				t.Error("could not match title")
				t.Errorf("want: %s", r.wantTitle)
				t.Errorf("got : %s", d)
			}
		})
	}
}