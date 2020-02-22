package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	log "go.uber.org/zap"
)

func TestRouter(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "tt non existing routes",
			fun:      testNonExistingRoute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			tt.fun(t)
		})
	}
}

func testNonExistingRoute(t *testing.T) {
	r := NewRouter(log.NewNop().Sugar())
	r.Get("/new", func(writer http.ResponseWriter, request *http.Request) {
	})
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/", nil))

	res := rec.Result()
	if http.StatusNotFound != res.StatusCode {
		t.Errorf("could not match status code")
	}
}
