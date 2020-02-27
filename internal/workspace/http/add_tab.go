package http

import (
	"encoding/json"
	"errors"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection/tab"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type addTabReq struct {
	Title       tab.Title       `json:"title"`
	Description tab.Description `json:"description"`
	Icon        tab.Icon        `json:"icon"`
	Link        tab.Link        `json:"link"`
}

func (r *addTabReq) UnmarshalJSON(data []byte) error {
	type clone addTabReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Title, err = tab.NewTitle(string(req.Title)); err != nil {
		return err
	}

	if r.Description, err = tab.NewDescription(string(req.Description)); err != nil {
		return err
	}

	if r.Icon, err = tab.NewIcon(string(req.Icon)); err != nil {
		return err
	}

	if r.Link, err = tab.NewLink(string(req.Link)); err != nil {
		return err
	}

	return nil
}

func addTab(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "addTab")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "addTab")

		id, err := repo.NextTabID(ctx)
		if err != nil {
			logger.With("error", err).Error("could not get next tab id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("id", id)
		var rb addTabReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, tab.ErrInvalidTitle) || errors.Is(err, tab.ErrInvalidDescription) || errors.Is(err, tab.ErrInvalidIcon) || errors.Is(err, tab.ErrInvalidLink):
			if _, err2 := w.Write([]byte(err.Error())); err2 != nil {
				logger.With("error", err, "error_2", err2).Error("could not write response")
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

		t := tab.New(id, rb.Title, rb.Description, rb.Icon, rb.Link)
		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		collID, ok := ctx.Value(collectionIDCtxKey).(collection.ID)
		if !ok {
			logger.Error("could not get collection id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !ws.AddTabs(collID, t) {
			logger.With("collection_id", collID).Info("could not find collection")
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
