package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/metabs/server/tab"
	"github.com/metabs/server/tab/collection"
	"github.com/metabs/server/tab/collection/workspace"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type patchTabReq = addTabReq

func putTab(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "putTab")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "putTab")

		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger = logger.With("workspace_id", ws.ID)

		collID, ok := ctx.Value(collectionIDCtxKey).(collection.ID)
		if !ok {
			logger.Error("could not get collection id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := tab.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Info("could not create tab id")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		var rb patchTabReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, tab.ErrInvalidTitle) || errors.Is(err, tab.ErrInvalidDescription) || errors.Is(err, tab.ErrInvalidIcon) || errors.Is(err, tab.ErrInvalidLink):
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

		t, ok := ws.FindTab(id, collID)
		if !ok {
			logger.Info("could not find tab")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		t.Update(rb.Title, rb.Description, rb.Icon, rb.Link)

		if !ws.UpdateTab(t, collID) {
			logger.Info("could not find collection or tab")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add tab")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(ws); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
