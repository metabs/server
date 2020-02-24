package tab

import (
	"errors"
	"testing"
)

func TestNewIcon(t *testing.T) {

	table := []struct {
		name     string
		raw      string
		wantIcon string
		wantErr  error
	}{
		{
			name:     "valid icon",
			raw:      "https://www.github.com/damianopetrungaro/image.png",
			wantIcon: "https://www.github.com/damianopetrungaro/image.png",
			wantErr:  nil,
		},
		{
			name:     "invalid icon",
			raw:      "a non url",
			wantIcon: "",
			wantErr:  ErrInvalidIcon,
		},
	}

	for _, r := range table {
		t.Run(r.name, func(t *testing.T) {
			i, err := NewIcon(r.raw)
			if !errors.Is(err, r.wantErr) {
				t.Error("could not match errors")
				t.Errorf("want: %s", r.wantErr)
				t.Errorf("got : %s", err)
			}

			if string(i) != r.wantIcon {
				t.Error("could not match icon")
				t.Errorf("want: %s", r.wantIcon)
				t.Errorf("got : %s", i)
			}
		})
	}
}
