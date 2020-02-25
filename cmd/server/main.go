package main

import (
	"context"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/workspace"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/log"

	database "github.com/unprogettosenzanomecheforseinizieremo/server/internal/db"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/http"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/probe"
)

func main() {
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	logger, err := log.New()
	if err != nil {
		panic(err)
	}

	go func() {
		<-kill
		defer cancel()
		signal.Stop(kill)
		close(kill)
		logger.Info("Stopping the application...")
	}()

	logger.Info("Application stopped.")
	db, err := database.New(ctx)
	if err != nil {
		logger.With("error", err).Fatal("Could not connect to the database.")
	}

	r := http.NewRouter(logger)
	r.Route("/", probe.NewRouter(db, logger))
	r.Route("/workspaces", workspace.NewRouter(&workspace.Repo{Client:db}, logger))
	srv := http.New(r)

	done := make(chan struct{}, 1)
	go func(done chan<- struct{}) {
		<-ctx.Done()

		logger.Info("Stopping the http...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.With("error", err).Fatal("Could not gracefully shutdown the http.")
		}

		// If you have any metrics or logs that need to be read before the shut down, remove the comment to the next 3 lines
		// logger.Info("Waiting metrics and logger to be read.")
		// <-ctx.Done()
		// logger.Info("Metrics and logger should be read.")

		logger.Info("Server stopped.")
		close(done)
	}(done)

	logger.Info("Server running...")
	if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
		logger.With("error", err).Fatal("Could not listen and serve.")
	}
	<-kill
}
