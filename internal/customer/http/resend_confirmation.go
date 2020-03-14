package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metabs/server/customer"
	"github.com/metabs/server/email"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type resendConfirmationReq struct {
	Email customer.Email `json:"email"`
}

func (r *resendConfirmationReq) UnmarshalJSON(data []byte) error {
	type clone resendConfirmationReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Email, err = customer.NewEmail(req.Email.String()); err != nil {
		return err
	}

	return nil
}

func resendConfirmation(repo customer.Repo, sender *email.Sender, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "resendConfirmation")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "resendConfirmation")

		var rb resendConfirmationReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, customer.ErrInvalidEmail):
			w.WriteHeader(http.StatusBadRequest)
			if _, err2 := w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error()))); err2 != nil {
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

		c, err := repo.GetByEmail(ctx, rb.Email)
		switch {
		case err != nil && !errors.Is(err, customer.ErrNotFound):
			logger.With("error", err).Error("could not get customer")
			w.WriteHeader(http.StatusInternalServerError)
			return

		case errors.Is(err, customer.ErrNotFound):
			logger.With("error", err).Error("customer not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if c.Status != customer.NotActivated {
			logger.With("error", err).Error("customer already confirmed email")
			w.WriteHeader(http.StatusBadRequest)
		}

		// TODO: when not a MVP use a queue
		if err := sender.SendActivationEmail(ctx, c.Email.String(), c.ID.String(), c.ActivateHash); err != nil {
			logger.With("error", err).Error("could not send email")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
