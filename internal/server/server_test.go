package server

import (
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "test server",
			fun:      testServerWithEnv,
		},
		{
			scenario: "test non existing routes",
			fun:      testServerEmptyEnv,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testServerWithEnv(t *testing.T) {
	address = ":80"
	readTimeout = "5"
	writeTimeout = "4"
	idleTimeout = "3"

	h := &http.ServeMux{}
	s := New(h)

	if s.Addr != address {
		t.Error("could not match addr")
		t.Errorf("got : %s", s.Addr)
		t.Errorf("want: %s", address)
	}
	if s.Handler != h {
		t.Error("could not match handler")
		t.Errorf("got : %v", s.Handler)
		t.Errorf("want: %v", h)
	}
	if s.ReadTimeout != time.Second*5 {
		t.Error("could not match readTimeout")
		t.Errorf("got : %s", s.ReadTimeout)
		t.Errorf("want: %s", time.Second*5)
	}
	if s.WriteTimeout != time.Second*4 {
		t.Error("could not match writeTimeout")
		t.Errorf("got : %s", s.WriteTimeout)
		t.Errorf("want: %s", time.Second*4)
	}
	if s.IdleTimeout != time.Second*3 {
		t.Error("could not match idleTimeout")
		t.Errorf("got : %s", s.IdleTimeout)
		t.Errorf("want: %s", time.Second*3)
	}
}

func testServerEmptyEnv(t *testing.T) {
	address = ":99"
	readTimeout = ""
	writeTimeout = ""
	idleTimeout = ""

	h := &http.ServeMux{}
	s := New(h)
	if s.Addr != address {
		t.Error("could not match addr")
		t.Errorf("got : %s", s.Addr)
		t.Errorf("want: %s", address)
	}
	if s.Handler != h {
		t.Error("could not match handler")
		t.Errorf("got : %v", s.Handler)
		t.Errorf("want: %v", h)
	}
	if s.ReadTimeout != 0 {
		t.Error("could not match readTimeout")
		t.Errorf("got : %s", s.ReadTimeout)
		t.Errorf("want: %d", 0)
	}
	if s.WriteTimeout != 0 {
		t.Error("could not match writeTimeout")
		t.Errorf("got : %s", s.WriteTimeout)
		t.Errorf("want: %d", 0)
	}
	if s.IdleTimeout != 0 {
		t.Error("could not match idleTimeout")
		t.Errorf("got : %s", s.IdleTimeout)
		t.Errorf("want: %d", 0)
	}
}
