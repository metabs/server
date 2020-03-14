package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/metabs/server/customer"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type changePasswordReq struct {
	Password customer.Password `json:"email"`
}

func (r *changePasswordReq) UnmarshalJSON(data []byte) error {
	type clone changePasswordReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Password, err = customer.NewPassword(req.Password.String()); err != nil {
		return err
	}

	return nil
}

func changePassword(repo customer.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "changePassword")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "changePassword")

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

		var rb changePasswordReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, customer.ErrInvalidPassword):
			w.WriteHeader(http.StatusUnauthorized)

			if _, err2 := w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`,err.Error()))); err2 != nil {
				logger.With("error", err, "error_2", err2).Error("could not write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			logger.With("error", err, "request", rb).Info("bad request")
			return

		case err != nil:
			logger.With("error", err).Error("could not decode request")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !c.ChangePassword(rb.Password, chi.URLParam(r, "hash")) {
			logger.With("error", err).Warn("could not change password")
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
