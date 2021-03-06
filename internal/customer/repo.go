package customer

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/metabs/server/customer"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const CollectionName = "customer"

type Repo struct {
	Client *firestore.Client
	Logger *zap.SugaredLogger
}

func (r *Repo) NextID(ctx context.Context) (customer.ID, error) {
	_, span := trace.StartSpan(ctx, "NextID")
	defer span.End()

	return customer.NewID(uuid.New().String())
}

func (r *Repo) Get(ctx context.Context, id customer.ID) (*customer.Customer, error) {
	ctx, span := trace.StartSpan(ctx, "Get")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", id, "action", "Get")

	doc, err := r.Client.Collection(CollectionName).Doc(id.String()).Get(ctx)
	switch {
	case status.Code(err) == codes.NotFound:
		logger.With("error", err).Info("document not found")
		return nil, fmt.Errorf("%w:%s", customer.ErrNotFound, err)

	case err != nil:
		logger.With("error", err).Errorf("could not find document")
		return nil, fmt.Errorf("%w:%s", customer.ErrRepoGet, err)
	}

	c := &customer.Customer{}
	if err := doc.DataTo(c); err != nil {
		logger.With("error", err).Errorf("could not map document")
		return nil, fmt.Errorf("%w:%s", customer.ErrRepoGet, err)
	}

	logger.Debug("gotten")
	return c, nil
}

func (r *Repo) GetByEmail(ctx context.Context, email customer.Email) (*customer.Customer, error) {
	ctx, span := trace.StartSpan(ctx, "GetByEmail")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "email", email, "action", "GetByEmail")

	switch docs, err := r.Client.Collection(CollectionName).Where("Email", "==", email.String()).Documents(ctx).GetAll(); {
	case err != nil:
		logger.With("error", err).Error("could not find document")
		return nil, fmt.Errorf("%w:%s", customer.ErrRepoGet, err)
	case len(docs) == 0:
		logger.With("error", err).Error("could not find document, none email found")
		return nil, customer.ErrNotFound
	case len(docs) != 1:
		logger.With("error", err).Error("could not return document, more than one email found")
		return nil, fmt.Errorf("%w:%s", customer.ErrRepoGet, "more than one doc found")
	default:
		c := &customer.Customer{}
		if err := docs[0].DataTo(c); err != nil {
			logger.With("error", err).Errorf("could not map document")
			return nil, fmt.Errorf("%w:%s", customer.ErrRepoGet, err)
		}

		logger.Debug("gotten")
		return c, nil
	}
}

func (r *Repo) Add(ctx context.Context, c *customer.Customer) error {
	ctx, span := trace.StartSpan(ctx, "Add")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", c.ID, "action", "Add")

	_, err := r.Client.Collection(CollectionName).Doc(c.ID.String()).Set(ctx, c)
	if err != nil {
		logger.With("error", err).Error("could not set document")
		return fmt.Errorf("%w:%s", customer.ErrRepoAdd, err)
	}

	logger.Debug("added")
	return nil
}

func (r *Repo) Delete(ctx context.Context, id customer.ID) error {
	ctx, span := trace.StartSpan(ctx, "Delete")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", id, "action", "Delete")

	if _, err := r.Client.Collection(CollectionName).Doc(id.String()).Delete(ctx); err != nil {
		logger.With("error", err).Error("could not delete document")
		return fmt.Errorf("%w:%s", customer.ErrRepoAdd, err)
	}

	return nil
}
