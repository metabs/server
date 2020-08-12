package http

import (
	"encoding/json"
	"github.com/metabs/server/tab/collection/workspace"

	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func list(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "list")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "list")

		cID, ok := ctx.Value(workspace.CustomerID("")).(workspace.CustomerID)
		if !ok {
			logger.Error("could not get customer id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("customer_id", cID)

		ws, err := repo.List(ctx, cID)
		if err != nil {
			logger.With("error", err).Error("could not list workspaces")
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
