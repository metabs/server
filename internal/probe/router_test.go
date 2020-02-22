package probe

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
)

func Test(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "test routes",
			fun:      testRoutes,
		},
		{
			scenario: "test liveness succed to ping server",
			fun:      testLivenessSucceedToPingServer,
		}, {
			scenario: "test liveness fail to ping server",
			fun:      testLivenessFailToPingServer,
		},
		{
			scenario: "test readiness succed to ping server",
			fun:      testReadinessSucceedToPingServer,
		}, {
			scenario: "test readiness fail to ping server",
			fun:      testReadinessFailToPingServer,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testRoutes(t *testing.T) {
	r := chi.NewRouter()
	r.Route("/", NewRouter(&sql.DB{}))
	if !r.Match(chi.NewRouteContext(), http.MethodGet, "/liveness") {
		t.Errorf("could not match /liveness path")
	}
	if !r.Match(chi.NewRouteContext(), http.MethodGet, "/readiness") {
		t.Errorf("could not match /readiness path")
	}
}

func testLivenessSucceedToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error("could not crete sql mock")
		t.Fatalf("got: %s", err)
	}
	db.SetMaxIdleConns(1)

	r := liveness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/liveness", nil))

	res := rec.Result()
	if http.StatusOK != res.StatusCode {
		t.Errorf("could not match status code")
		t.Errorf("want: %d", http.StatusOK)
		t.Errorf("got : %d", res.StatusCode)
	}

	got, err := ioutil.ReadAll(res.Body)
	want := `["Server is live"]`
	if want != string(got) {
		t.Errorf("could not match response body")
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", got)
	}
}

func testLivenessFailToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error("could not crete sql mock")
		t.Fatalf("got: %s", err)
	}
	db.SetMaxIdleConns(0)

	r := liveness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/liveness", nil))

	res := rec.Result()
	if http.StatusInternalServerError != res.StatusCode {
		t.Errorf("could not match status code")
		t.Errorf("want: %d", http.StatusInternalServerError)
		t.Errorf("got : %d", res.StatusCode)
	}

	got, err := ioutil.ReadAll(res.Body)
	want := `["Database is not alive"]`
	if want != string(got) {
		t.Errorf("could not match response body")
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", got)
	}
}

func testReadinessSucceedToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error("could not crete sql mock")
		t.Fatalf("got: %s", err)
	}
	db.SetMaxIdleConns(1)

	r := readiness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/readiness", nil))

	res := rec.Result()
	if http.StatusOK != res.StatusCode {
		t.Errorf("could not match status code")
		t.Errorf("want: %d", http.StatusOK)
		t.Errorf("got : %d", res.StatusCode)
	}

	got, err := ioutil.ReadAll(res.Body)
	want := `["Server is ready"]`
	if want != string(got) {
		t.Errorf("could not match response body")
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", got)
	}

}

func testReadinessFailToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error("could not crete sql mock")
		t.Fatalf("got: %s", err)
	}
	db.SetMaxIdleConns(0)

	r := readiness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/readiness", nil))

	res := rec.Result()
	if http.StatusInternalServerError != res.StatusCode {
		t.Errorf("could not match status code")
		t.Errorf("want: %d", http.StatusInternalServerError)
		t.Errorf("got : %d", res.StatusCode)
	}

	got, err := ioutil.ReadAll(res.Body)
	want := `["Database is not alive"]`
	if want != string(got) {
		t.Errorf("could not match response body")
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", got)
	}
}
