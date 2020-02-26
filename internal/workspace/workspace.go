package workspace

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"go.uber.org/zap"
	"net/http"
)

func add(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "create workspace")
		ctx := r.Context()

		id, err := repo.NextID(ctx)
		if err != nil {
			logger.With("error", err).Error("could not get next ID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("id", id)

		cID, ok := ctx.Value(workspace.CustomerID("")).(workspace.CustomerID)
		if !ok {
			logger.Error("could not get customer ID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("customer_id", cID)

		type requestBody struct {
			Name workspace.Name `json:"name,string"`
		}
		var rb requestBody
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, workspace.ErrInvalidName):
			if _, wErr := w.Write([]byte(err.Error())); wErr != nil {
				logger.With("error", err, "error_2", wErr).Error("could not write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		case err != nil:
			logger.With("error", err).Error("could not decode request")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("name", rb.Name)

		ws := workspace.New(id, rb.Name, cID)
		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(ws); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func list(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "list CollectionName")
		ctx := r.Context()

		cID, ok := ctx.Value(workspace.CustomerID("")).(workspace.CustomerID)
		if !ok {
			logger.Error("could not get customer ID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("customer_id", cID)

		ws, err := repo.List(ctx, cID)
		if err != nil {
			logger.With("error", err).Error("could not list workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(ws); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func delete(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "delete CollectionName")
		ctx := r.Context()

		id, err := workspace.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Error("could not validate uuid")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		switch err := repo.Delete(ctx, id); {
		case err == workspace.ErrNotFound:
			logger.With("error", err).Info("not found")
			w.WriteHeader(http.StatusNotFound)
			return
		case err != nil:
			logger.With("error", err).Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func get(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "get workspace")
		ctx := r.Context()

		id, err := workspace.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Error("could not validate uuid")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		ws, err := repo.Get(ctx, id)
		switch {
		case err == workspace.ErrNotFound:
			logger.With("error", err).Info("not found")
			w.WriteHeader(http.StatusNotFound)
			return
		case err != nil:
			logger.With("error", err).Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(ws); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func patch(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "patch workspace")
		ctx := r.Context()

		id, err := workspace.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Error("could not validate uuid")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logger = logger.With("id", id)

		ws, err := repo.Get(ctx, id)
		switch {
		case err == workspace.ErrNotFound:
			logger.With("error", err).Info("not found")
			w.WriteHeader(http.StatusNotFound)
			return
		case err != nil:
			logger.With("error", err).Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		type requestBody struct {
			Name workspace.Name `json:"name,string"`
		}
		var rb requestBody
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, workspace.ErrInvalidName):
			if _, wErr := w.Write([]byte(err.Error())); wErr != nil {
				logger.With("error", err, "error_2", wErr).Error("could not write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		case err != nil:
			logger.With("error", err).Error("could not decode request")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("name", rb.Name)

		ws.Rename(rb.Name)

		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(ws); err != nil {
			logger.With("error", err).Error("could not write response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
