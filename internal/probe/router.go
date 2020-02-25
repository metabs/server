package probe

import (
	"cloud.google.com/go/firestore"
	"go.uber.org/zap"
	"net/http"

	"github.com/go-chi/chi"
)

// NewRouter Return a function to use with an existing router
func NewRouter(db *firestore.Client, log *zap.SugaredLogger) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/liveness", liveness(db, log))
		r.Get("/readiness", readiness(db, log))
	}
}

func liveness(db *firestore.Client, log *zap.SugaredLogger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger := log.With("action", "liveness")
		if _, err := db.Collection("Ping").Doc("0").Get(req.Context()); err != nil {
			logger.With("error", err).Error("could not commit empty transaction")
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(`["Database is not alive"]`))
			return
		}

		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte(`["Server is live"]`))
	}
}

func readiness(db *firestore.Client, log *zap.SugaredLogger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger := log.With("action", "readiness")
		if _, err := db.Collection("Ping").Doc("0").Get(req.Context()); err != nil {
			logger.With("error", err).Error("could not commit empty transaction")
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(`["Database is not alive"]`))
			return
		}

		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte(`["Server is ready"]`))
	}
}
