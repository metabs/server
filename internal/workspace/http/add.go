package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metabs/server/workspace"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type addReq struct {
	Name workspace.Name `json:"name"`
}

func (r *addReq) UnmarshalJSON(data []byte) error {
	type clone addReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Name, err = workspace.NewName(req.Name.String()); err != nil {
		return err
	}

	return nil
}

func add(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "add")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "add")

		id, err := repo.NextID(ctx)
		if err != nil {
			logger.With("error", err).Error("could not get next id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("id", id)

		cID, ok := ctx.Value(workspace.CustomerID("")).(workspace.CustomerID)
		if !ok {
			logger.Error("could not get customer ID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("customer_id", cID)

		var rb addReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, workspace.ErrInvalidName):
			if _, err2 := w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`,err.Error()))); err2 != nil {
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

		ws := workspace.New(id, rb.Name, cID)
		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add workspace")
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
