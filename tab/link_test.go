package tab

import (
	"errors"
	"testing"
)

func TestNewLink(t *testing.T) {

	table := []struct {
		name     string
		raw      string
		wantLink string
		wantErr  error
	}{
		{
			name:     "valid title",
			raw:      "https://www.github.com/damianopetrungaro/image.png",
			wantLink: "https://www.github.com/damianopetrungaro/image.png",
			wantErr:  nil,
		},
		{
			name:     "invalid title",
			raw:      "a non url",
			wantLink: "",
			wantErr:  ErrInvalidLink,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			l, err := NewLink(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}

			if l.String() != r.wantLink {
				t.Error("could not match title")
				t.Errorf("want: %s", r.wantLink)
				t.Errorf("got : %s", l)
			}
		})
	}
}
