package probe

import (
	"cloud.google.com/go/firestore"
	"go.opencensus.io/trace"
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
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "liveness")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "liveness")

		if _, err := db.Collection("Ping").Doc("0").Get(ctx); err != nil {
			logger.With("error", err).Error("could not commit empty transaction")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`["Database is not alive"]`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["Server is live"]`))
	}
}

func readiness(db *firestore.Client, log *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "readiness")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "readiness")

		if _, err := db.Collection("Ping").Doc("0").Get(ctx); err != nil {
			logger.With("error", err).Error("could not commit empty transaction")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`["Database is not alive"]`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["Server is ready"]`))
	}
}
