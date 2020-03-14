package http

import (
	"github.com/metabs/server/customer"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func delete(repo customer.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "delete")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "delete")
		ws, ok := ctx.Value(customerCtxKey).(*customer.Customer)
		if !ok {
			logger.Error("could not get customer")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := repo.Delete(ctx, ws.ID); err != nil {
			logger.With("error", err).Error("could not delete customer")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
