package http

import (
	"github.com/go-chi/chi"
	"github.com/metabs/server/tab"
	"github.com/metabs/server/tab/collection"
	"github.com/metabs/server/tab/collection/workspace"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func deleteTab(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "deleteTab")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "deleteTab")

		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger = logger.With("workspace_id", ws.ID)

		id, err := tab.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Info("could not create tab id")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		collID, ok := ctx.Value(collectionIDCtxKey).(collection.ID)
		if !ok {
			logger.Error("could not get collection id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger = logger.With("collection_id", collID)

		if !ws.RemoveTab(id, collID) {
			logger.Info("could not find collection or tab")
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
