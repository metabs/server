package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/metabs/server/tab/collection"
	"github.com/metabs/server/tab/collection/workspace"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type patchCollectionReq = addCollectionReq

func patchCollection(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "patchCollection")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "patchCollection")

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

		var rb patchCollectionReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, collection.ErrInvalidName):
			if _, wErr := w.Write([]byte(err.Error())); wErr != nil {
				logger.With("error", err, "error_2", wErr).Error("could not write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			logger.With("error", err, "request", rb).Info("bad request")
			w.WriteHeader(http.StatusBadRequest)
			return

		case err != nil:
			logger.With("error", err).Error("could not decode request")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !ws.RenameCollection(id, rb.Name) {
			logger.With("error", err).Error("could not remove collection, not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add collection")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
