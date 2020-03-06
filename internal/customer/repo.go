package customer

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
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

	doc, err := r.findByID(ctx, id)
	if err != nil {
		logger.With("error", err).Error("could not find document")
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

	doc, err := r.Client.Collection(CollectionName).Doc(email.String()).Get(ctx)
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

func (r *Repo) Add(ctx context.Context, c *customer.Customer) error {
	ctx, span := trace.StartSpan(ctx, "Add")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", c.ID, "action", "Add")

	_, err := r.Client.Collection(CollectionName).Doc(c.Email.String()).Set(ctx, c)
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

	doc, err := r.findByID(ctx, id)
	if err != nil {
		logger.With("error", err).Error("could not find document")
		return fmt.Errorf("%w:%s", customer.ErrRepoDelete, err)
	}

	email, ok := doc.Data()["Email"].(string)
	if !ok {
		logger.With("error", err).Error("could not delete document, more than one id found")
		return fmt.Errorf("%w:%s", customer.ErrRepoDelete, "could not transform email")
	}

	if _, err := r.Client.Collection(CollectionName).Doc(email).Delete(ctx); err != nil {
		logger.With("error", err).Error("could not delete document")
		return fmt.Errorf("%w:%s", customer.ErrRepoAdd, err)
	}

	return nil
}

func (r *Repo) findByID(ctx context.Context, id customer.ID) (*firestore.DocumentSnapshot, error) {
	ctx, span := trace.StartSpan(ctx, "findByID")
	defer span.End()

	logger := r.Logger.With("trace_id", span.SpanContext().TraceID.String(), "id", id, "action", "findByID")

	switch docs, err := r.Client.Collection(CollectionName).Where("ID", "==", id.String()).Documents(ctx).GetAll(); {
	case err != nil:
		logger.With("error", err).Error("could not find document")
		return nil, err
	case len(docs) == 0:
		logger.With("error", err).Error("could not find document, none id found")
		return nil, errors.New("no doc found")
	case len(docs) != 1:
		logger.With("error", err).Error("could not return document, more than one id found")
		return nil, errors.New("more than one doc found")
	default:
		return docs[0], nil
	}
}
