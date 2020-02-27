package http

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/metabs/server/customer"
	"github.com/metabs/server/internal/jwt"
	"github.com/metabs/server/workspace"
	"github.com/metabs/server/workspace/collection"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

type (
	workspaceCtx    int
	collectionIDCtx int
)

var workspaceCtxKey workspaceCtx = 0
var collectionIDCtxKey collectionIDCtx = 0

// NewRouter Return a function to use with an existing router
func NewRouter(repo workspace.Repo, sv *jwt.SignerVerifier, log *zap.SugaredLogger) func(r chi.Router) {
	return func(r chi.Router) {
		r = r.With(middleware.Timeout(time.Millisecond*150), jwtMiddleware(sv, log))
		r.Post("/", add(repo, log))
		r.Get("/", list(repo, log))
		r.With(workspaceMiddleware(repo, "id", log)).Delete("/{id}", delete(repo, log))
		r.With(workspaceMiddleware(repo, "id", log)).Get("/{id}", get(repo, log))
		r.With(workspaceMiddleware(repo, "id", log)).Patch("/{id}", patch(repo, log))

		r.With(workspaceMiddleware(repo, "workspace_id", log)).Route("/{workspace_id}/collections", func(r chi.Router) {
			r.Post("/", addCollection(repo, log))
			r.Patch("/{id}", patchCollection(repo, log))
			r.Delete("/{id}", deleteCollection(repo, log))

			r.With(collectionMiddleware(log)).Route("/{collection_id}/tabs", func(r chi.Router) {
				r.Post("/", addTab(repo, log))
				r.Put("/{id}", putTab(repo, log))
				r.Delete("/{id}", deleteTab(repo, log))
			})
		})
	}
}

func jwtMiddleware(sv *jwt.SignerVerifier, log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx, span := trace.StartSpan(r.Context(), "jwtMiddleware")
			defer span.End()

			logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "jwtMiddleware")

			h := r.Header.Get("Authorization")
			if !strings.HasPrefix(h, "Bearer ") {
				logger.Warn("could not verify non bearer token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			h = strings.TrimPrefix(h, "Bearer ")
			cID, cStatus, err := sv.Verify(h)
			if err != nil {
				logger.With("error", err).Warn("could not verify jwt token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if cStatus != customer.Activated {
				logger.With("error", err).Warn("customer not active")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			logger.With("customer_id", cID).Debug("authorized")

			ctx = context.WithValue(ctx, workspace.CustomerID(""), workspace.CustomerID(cID.String()))
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func workspaceMiddleware(repo workspace.Repo, k string, log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx, span := trace.StartSpan(r.Context(), "workspaceMiddleware")
			defer span.End()

			logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "workspaceMiddleware", "key", k)

			id, err := workspace.NewID(chi.URLParam(r, k))
			if err != nil {
				logger.With("error", err).Error("could not create workspace id")
				w.WriteHeader(http.StatusNotFound)
				return
			}

			logger = logger.With("id", id)

			cID, ok := ctx.Value(workspace.CustomerID("")).(workspace.CustomerID)
			if !ok {
				logger.Error("could not get customer id")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			logger = logger.With("customer_id", cID)

			ws, err := repo.Get(ctx, id)
			switch {
			case errors.Is(err, workspace.ErrNotFound):
				logger.With("error", err).Info("not found")
				w.WriteHeader(http.StatusNotFound)
				return

			case err != nil:
				logger.With("error", err).Error("could not get workspace")
				w.WriteHeader(http.StatusInternalServerError)
				return

			case ws.CustomerID != cID:
				logger.With().Warn("could not access resource")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(ctx, workspaceCtxKey, ws)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func collectionMiddleware(log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx, span := trace.StartSpan(r.Context(), "collectionMiddleware")
			defer span.End()

			logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "collectionMiddleware")

			id, err := collection.NewID(chi.URLParam(r, "collection_id"))
			if err != nil {
				logger.With("error", err).Error("could not create collection id")
				w.WriteHeader(http.StatusNotFound)
				return
			}

			ctx = context.WithValue(ctx, collectionIDCtxKey, id)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
