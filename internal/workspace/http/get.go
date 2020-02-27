package http

import (
	"encoding/json"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func get(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "get")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "get")

		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger = logger.With("workspace_id", ws.ID)


		if err := json.NewEncoder(w).Encode(ws); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
