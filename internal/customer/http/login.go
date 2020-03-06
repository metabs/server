package http

import (
	"encoding/json"
	"errors"
	"github.com/metabs/server/customer"
	"github.com/metabs/server/internal/jwt"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

type loginReq struct {
	Email    customer.Email `json:"email"`
	Password string         `json:"password"`
}

func (r *loginReq) UnmarshalJSON(data []byte) error {
	type clone loginReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Email, err = customer.NewEmail(req.Email.String()); err != nil {
		return err
	}
	if _, err = customer.NewPassword(req.Password); err != nil {
		return err
	}

	r.Password = req.Password
	return nil
}

func login(repo customer.Repo, sv *jwt.SignerVerifier, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "login")
		defer span.End()

		logger := log.With("trace_id", span.SpanContext().TraceID.String(), "action", "login")

		var rb loginReq
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, customer.ErrInvalidEmail) || errors.Is(err, customer.ErrInvalidPassword):
			w.WriteHeader(http.StatusUnauthorized)

			if _, err2 := w.Write([]byte(err.Error())); err2 != nil {
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

		logger = logger.With("email", rb.Email)
		c, err := repo.GetByEmail(ctx, rb.Email)
		if err != nil {
			logger.With("error", err).Warn("could not find customer")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !c.Password.Compare(rb.Password) {
			logger.With("error", err).Warn("wrong password")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !c.IsActive() {
			logger.With("error", err).Info("not active")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(c); err != nil {
				logger.With("error", err).Error("could not write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		token, err := sv.Sign(c.ID, c.Email, c.Status, c.Created, c.Activated, c.Updated, )
		if err != nil {
			logger.With("error", err).Error("could not sign customer")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(token); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
