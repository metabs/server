package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/metabs/server/customer"
	"github.com/metabs/server/email"
	"github.com/metabs/server/internal/jwt"
	"go.uber.org/zap"
	"time"
)

type customerCtx int

var customerCtxKey customerCtx = 0

// NewRouter Return a function to use with an existing router
func NewRouter(repo customer.Repo, sv *jwt.SignerVerifier, sender *email.Sender, log *zap.SugaredLogger) func(r chi.Router) {
	return func(r chi.Router) {
		r = r.With(middleware.Timeout(time.Second * 2))
		r.Post("/", signUp(repo, sender, log))
		r.Post("/resend", resendConfirmation(repo, sender, log))
		r.Patch("/{id}/activate/{hash}", activate(repo, log))
		r.Patch("/{id}/password/{hash}", changePassword(repo, log))
		r.Patch("/{id}/email/{hash}", changeEmail(repo, sender, log))
		r.Post("/{id}/password", resetPassword(repo, sender, log))
		r.Post("/{id}/email", newEmail(repo, sender, log))
		r.Post("/login", login(repo, sv, log))
		r.Delete("/non_working_api", delete(repo, log))
	}
}
