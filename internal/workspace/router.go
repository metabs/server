package workspace

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"go.uber.org/zap"
	"net/http"
)

type workspaceCtxKey string

// NewRouter Return a function to use with an existing router
func NewRouter(repo workspace.Repo, collRepo collection.Repo, log *zap.SugaredLogger) func(r chi.Router) {
	return func(r chi.Router) {
		r = r.With(jwtMiddleware())
		r.Post("/", add(repo, log))
		r.Get("/", list(repo, log))
		r.Delete("/{id}", delete(repo, log))
		r.Get("/{id}", get(repo, log))
		r.Patch("/{id}", patch(repo, log))

		r.Route("/{workspace_id}/collections", func(r chi.Router) {
			r = r.With(workspaceMiddleware(repo, log))
			r.Post("/", addCollection(repo, collRepo, log))
			r.Patch("/{id}", patchCollection(repo, log))
			r.Delete("/{id}", deleteCollection(repo, log))

			r.Route("/{collection_id}/tabs/", func(r chi.Router) {
				// r.Post("/", addCollection(repo, collRepo, log))
				// r.Patch("/{id}", patchCollection(repo, log))
				// r.Delete("/{id}", deleteCollection(repo, log))
			})
		})
	}
}

func jwtMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// Silly middleware to power fake customer id
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), workspace.CustomerID(""), workspace.CustomerID("1ae3a55d-2c69-4679-808e-1c7772405281")))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func workspaceMiddleware(repo workspace.Repo, log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// refactor: get and this are duplicated :)
			logger := log.With("action", "middleware workspace")
			ctx := r.Context()

			id, err := workspace.NewID(chi.URLParam(r, "workspace_id"))
			if err != nil {
				logger.With("error", err).Error("could not validate uuid")
				w.WriteHeader(http.StatusNotFound)
				return
			}

			logger = logger.With("workspace_id", id)

			ws, err := repo.Get(ctx, id)
			switch {
			case err == workspace.ErrNotFound:
				logger.With("error", err).Info("not found")
				w.WriteHeader(http.StatusBadRequest)
				return
			case err != nil:
				logger.With("error", err).Error("could not get workspace")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = context.WithValue(ctx, workspaceCtxKey(""), ws)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
