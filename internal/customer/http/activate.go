package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/metabs/server/customer"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

func activate(repo customer.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "activate")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "activate")

		id, err := customer.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Info("could not create id")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)
		c, err := repo.Get(ctx, id)
		if err != nil {
			logger.With("error", err).Warn("could not find customer")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if !c.Activate(chi.URLParam(r, "hash")) {
			logger.With("error", err).Warn("could not activate customer")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := repo.Add(ctx, c); err != nil {
			logger.With("error", err).Error("could not add customer")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(c); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
