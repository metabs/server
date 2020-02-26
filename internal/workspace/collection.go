package workspace

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"go.uber.org/zap"
	"net/http"
)

func addCollection(repo workspace.Repo, collRepo collection.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "create collection")
		ctx := r.Context()

		id, err := collRepo.NextID(ctx)
		if err != nil {
			logger.With("error", err).Error("could not get next ID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger = logger.With("id", id)

		type requestBody struct {
			Name collection.Name `json:"name,string"`
		}
		var rb requestBody
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, collection.ErrInvalidName):
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
		col := collection.New(id, rb.Name)

		ws, ok := ctx.Value(workspaceCtxKey("")).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ws.AddCollections(col)
		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add collection")
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

func deleteCollection(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "delete CollectionName")
		ctx := r.Context()

		id, err := collection.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Error("could not validate uuid")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		ws, ok := ctx.Value(workspaceCtxKey("")).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !ws.RemoveCollection(id) {
			return
		}
		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add collection")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func patchCollection(repo workspace.Repo, log *zap.SugaredLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With("action", "patch CollectionName")
		ctx := r.Context()

		id, err := collection.NewID(chi.URLParam(r, "id"))
		if err != nil {
			logger.With("error", err).Error("could not validate uuid")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger = logger.With("id", id)

		ws, ok := ctx.Value(workspaceCtxKey("")).(*workspace.Workspace)
		if !ok {
			logger.Error("could not get workspace")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		type requestBody struct {
			Name collection.Name `json:"name,string"`
		}
		var rb requestBody
		switch err := json.NewDecoder(r.Body).Decode(&rb); {
		case errors.Is(err, collection.ErrInvalidName):
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

		if !ws.RenameCollection(id, rb.Name) {
			logger.With("error", err).Error("could not remove collection, not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := repo.Add(ctx, ws); err != nil {
			logger.With("error", err).Error("could not add collection")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
