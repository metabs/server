package customer

import (
	"testing"
)

func TestNew(t *testing.T) {
	id := ID("a")
	email := Email("a")
	pwd := Password("xx")
	customer := New(id, email, pwd)

	if customer.ID != id {
		t.Error("could not match id")
		t.Errorf("want: %s", id)
		t.Errorf("got : %s", customer.ID)
	}

	if customer.Email != email {
		t.Error("could not match email")
		t.Errorf("want: %s", email)
		t.Errorf("got : %s", customer.Email)
	}

	if customer.Password != pwd {
		t.Error("could not match password")
		t.Errorf("want: %s", pwd)
		t.Errorf("got : %s", customer.Password)
	}

	if customer.Created.IsZero() {
		t.Error("could not find a created time")
		t.Errorf("got : %s", customer.Created)
	}
}
