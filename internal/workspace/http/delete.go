package http

import (
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func delete(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "delete")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "delete")

		ws, ok := ctx.Value(workspaceCtxKey).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := repo.Delete(ctx, ws.ID); err != nil {
			logger.With("error", err).Error("could not delete workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
