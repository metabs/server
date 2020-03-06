package http

import (
	"github.com/go-chi/chi"
	"github.com/metabs/server/workspace"
	"github.com/metabs/server/workspace/collection"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func deleteCollection(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "deleteCollection")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "deleteCollection")

		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger = logger.With("workspace_id", ws.ID)


		id, err := collection.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Info("could not create collection id")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		if !ws.RemoveCollection(id) {
			logger.Info("could not delete non existing workspace")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
