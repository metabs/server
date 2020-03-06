package http

import (
	"encoding/json"
	"errors"
	"github.com/metabs/server/workspace"
	"github.com/metabs/server/workspace/collection"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type addCollectionReq struct {
	Name collection.Name `json:"name"`
}

func (r *addCollectionReq) UnmarshalJSON(data []byte) error {
	type clone addCollectionReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	var err error
	if r.Name, err = collection.NewName(req.Name.String()); err != nil {
		return err
	}

	return nil
}

func addCollection(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "addCollection")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "addCollection")

		id, err := repo.NextCollectionID(ctx)
		if err != nil {
			logger.With("error", err).Error("could not get next collection id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("id", id)

		var rb addCollectionReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, collection.ErrInvalidName):
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

		col := collection.New(id, rb.Name)

		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ws.AddCollections(col)
		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add collection")
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