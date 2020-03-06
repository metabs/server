package workspace

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/metabs/server/workspace"
	"github.com/metabs/server/workspace/collection"
	"github.com/metabs/server/workspace/collection/tab"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const CollectionName = "workspace"

type Repo struct {
	Client *firestore.Client
	Logger *zap.SugaredLogger
}

func (r *Repo) NextID(ctx context.Context) (workspace.ID, error) {
	_, span := trace.StartSpan(ctx, "NextID")
	defer span.End()

	return workspace.NewID(uuid.New().String())
}

func (r *Repo) NextCollectionID(ctx context.Context) (collection.ID, error) {
	_, span := trace.StartSpan(ctx, "NextCollectionID")
	defer span.End()

	return collection.NewID(uuid.New().String())
}

func (r *Repo) NextTabID(ctx context.Context) (tab.ID, error) {
	_, span := trace.StartSpan(ctx, "NextTabID")
	defer span.End()

	return tab.NewID(uuid.New().String())
}

func (r *Repo) List(ctx context.Context, cID workspace.CustomerID) ([]*workspace.Workspace, error) {
	ctx, span := trace.StartSpan(ctx, "List")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "customer_id", cID, "action", "List")

	docs := r.Client.Collection(CollectionName).Where("CustomerID", "==", cID.String()).Documents(ctx)
	var wss = make([]*workspace.Workspace, 0)
	for {
		doc, err := docs.Next()
		switch {
		case err == iterator.Done:
			logger.Debug("listed")
			return wss, nil

		case err != nil:
			logger.With("error", err).Error("could not map next document")
			return nil, fmt.Errorf("%w:%s", workspace.ErrRepoList, err)

		default:

			ws := &workspace.Workspace{Collections: make([]*collection.Collection, 0)}
			if err := doc.DataTo(ws); err != nil {
				logger.With("error", err).Error("could not map document")
				return wss, fmt.Errorf("%w:%s", workspace.ErrRepoList, err)
			}
			wss = append(wss, ws)
		}
	}
}

func (r *Repo) Get(ctx context.Context, id workspace.ID) (*workspace.Workspace, error) {
	ctx, span := trace.StartSpan(ctx, "Get")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", id, "action", "Get")

	doc, err := r.Client.Collection(CollectionName).Doc(id.String()).Get(ctx)
	switch {
	case status.Code(err) == codes.NotFound:
		logger.With("error", err).Info("document not found")
		return nil, fmt.Errorf("%w:%s", workspace.ErrNotFound, err)

	case err != nil:
		logger.With("error", err).Errorf("could not find document")
		return nil, fmt.Errorf("%w:%s", workspace.ErrRepoGet, err)
	}

	ws := &workspace.Workspace{Collections: make([]*collection.Collection, 0)}
	if err := doc.DataTo(ws); err != nil {
		logger.With("error", err).Errorf("could not map document")
		return nil, fmt.Errorf("%w:%s", workspace.ErrRepoGet, err)
	}

	logger.Debug("gotten")
	return ws, nil
}

func (r *Repo) Add(ctx context.Context, ws *workspace.Workspace) error {
	ctx, span := trace.StartSpan(ctx, "Add")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", ws.ID, "action", "Add")

	_, err := r.Client.Collection(CollectionName).Doc(ws.ID.String()).Set(ctx, ws)
	if err != nil {
		logger.With("error", err).Error("could not set document")
		return fmt.Errorf("%w:%s", workspace.ErrRepoAdd, err)
	}

	logger.Debug("added")
	return nil
}

func (r *Repo) Delete(ctx context.Context, id workspace.ID) error {
	ctx, span := trace.StartSpan(ctx, "Delete")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", id, "action", "Delete")
	
	if _, err := r.Client.Collection(CollectionName).Doc(id.String()).Delete(ctx); err != nil {
		logger.With("error", err).Error("could not delete document")
		return fmt.Errorf("%w:%s", workspace.ErrRepoDelete, err)
	}

	return nil
}
