package http

import (
	"go.opencensus.io/trace"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const traceIDHeader = "X-Trace-ID"

// NewRouter Return a new basic router with some handy middleware
func NewRouter(log *zap.SugaredLogger) chi.Router {
	r := chi.NewRouter()

	r.Use(
		corsMiddleware(),
		profilingMiddleware(log),
		middleware.Timeout(time.Second*10),
		middleware.SetHeader("Content-Type", "application/json"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	return r
}

func profilingMiddleware(log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := trace.StartSpan(r.Context(), "profilingMiddleware", trace.WithSpanKind(trace.SpanKindServer))
			defer span.End()

			r = r.WithContext(ctx)
			rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				start := time.Now()
				log.With(
					"status_code", rw.Status(),
					"http_verb", r.Method,
					"bytes", rw.BytesWritten(),
					"latency", time.Since(start).Seconds(),
					"uri", r.URL.String(),
					"trace_id", r.Header.Get(traceIDHeader)).
					Info("router: api call done")
			}()
			next.ServeHTTP(rw, r)
		})
	}
}

func corsMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-Requested-With, Authorization")
			if r.Method == http.MethodOptions {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
