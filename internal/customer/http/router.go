package http

import (
	"github.com/go-chi/chi"
	"github.com/unprogettosenzanomecheforseinizieremo/server/customer"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/jwt"
	"go.uber.org/zap"
)

type customerCtx int

var customerCtxKey customerCtx = 0

// NewRouter Return a function to use with an existing router
func NewRouter(repo customer.Repo, sv *jwt.SignerVerifier, log *zap.SugaredLogger) func(r chi.Router) {
	return func(r chi.Router) {
		//r = r.With(jwtMiddleware(log))
		r.Post("/", signUp(repo, log))
		r.Patch("/{id}/activate/{hash}", activate(repo, log))
		r.Post("/login", login(repo, sv, log))
		r.Delete("/{id}", delete(repo, log))
	}
}
