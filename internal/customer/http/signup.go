package http

import (
	"encoding/json"
	"errors"
	"github.com/unprogettosenzanomecheforseinizieremo/server/customer"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type signUpReq struct {
	Email    customer.Email    `json:"email"`
	Password customer.Password `json:"password"`
}

func (r *signUpReq) UnmarshalJSON(data []byte) error {
	type clone signUpReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Email, err = customer.NewEmail(req.Email.String()); err != nil {
		return err
	}
	if r.Password, err = customer.NewPassword(req.Password.String()); err != nil {
		return err
	}

	return nil
}

func signUp(repo customer.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "sign up")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "sign up")

		id, err := repo.NextID(ctx)
		if err != nil {
			logger.With("error", err).Error("could not get next id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("id", id)

		var rb signUpReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, customer.ErrInvalidEmail) || errors.Is(err, customer.ErrInvalidPassword):
			w.WriteHeader(http.StatusBadRequest)
			if _, err2 := w.Write([]byte(err.Error())); err2 != nil {
				logger.With("error", err, "error_2", err2).Error("could not write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			logger.With("error", err, "email", rb.Email).Info("bad request")
			return

		case err != nil:
			logger.With("error", err).Error("could not decode request")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch _, err := repo.GetByEmail(ctx, rb.Email); {
		case err != nil && !errors.Is(err, customer.ErrNotFound):
			logger.With("error", err).Error("could not get customer")
			w.WriteHeader(http.StatusInternalServerError)
			return

		case err == nil:
			logger.With("error", err).Error("customer already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		c := customer.New(id, rb.Email, rb.Password)
		if err := repo.Add(ctx, c); err != nil {
			logger.With("error", err).Error("could not add customer")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: when not a MVP use a queue
		// send activation email here

		if err := json.NewEncoder(w).Encode(c); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
