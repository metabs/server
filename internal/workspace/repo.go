package workspace

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace"
	"github.com/unprogettosenzanomecheforseinizieremo/server/workspace/collection"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const CollectionName = "workspace"

type Repo struct {
	Client *firestore.Client
}

type CollectionRepo struct {
}

func (r *CollectionRepo) NextID(_ context.Context) (collection.ID, error) {
	return collection.NewID(uuid.New().String())
}

func (r *Repo) NextID(_ context.Context) (workspace.ID, error) {
	return workspace.NewID(uuid.New().String())
}

func (r *Repo) List(ctx context.Context, cID workspace.CustomerID) ([]*workspace.Workspace, error) {
	docs := r.Client.Collection(CollectionName).Where("CustomerID", "==", string(cID)).Documents(ctx)

	var wss = make([]*workspace.Workspace, 0)
	for {
		doc, err := docs.Next()
		switch {
		case err == iterator.Done:
			return wss, nil
		case err != nil:
			return nil, fmt.Errorf("%w:%s", workspace.ErrRepoList, err)
		default:
			ws := &workspace.Workspace{Collections: make([]*collection.Collection, 0)}
			if err := doc.DataTo(ws); err != nil {
				return wss, nil
			}
			wss = append(wss, ws)
		}
	}
}

func (r *Repo) Get(ctx context.Context, id workspace.ID) (*workspace.Workspace, error) {
	doc, err := r.Client.Collection(CollectionName).Doc(string(id)).Get(ctx)
	switch {
	case status.Code(err) == codes.NotFound:
		return nil, workspace.ErrNotFound
	case err != nil:
		return nil, fmt.Errorf("%w:%s", workspace.ErrRepoGet, err)
	}

	ws := &workspace.Workspace{Collections: make([]*collection.Collection, 0)}
	if err := doc.DataTo(ws); err != nil {
		return nil, fmt.Errorf("%w:%s", workspace.ErrRepoGet, err)
	}

	return ws, nil
}

func (r *Repo) Add(ctx context.Context, ws *workspace.Workspace) error {
	_, err := r.Client.Collection(CollectionName).Doc(string(ws.ID)).Set(ctx, ws)
	if err != nil {
		return fmt.Errorf("%w:%s", workspace.ErrRepoAdd, err)
	}

	return nil
}

func (r *Repo) Delete(ctx context.Context, id workspace.ID) error {
	_, err := r.Client.Collection(CollectionName).Doc(string(id)).Get(ctx)
	switch {
	case status.Code(err) == codes.NotFound:
		return workspace.ErrNotFound
	case err != nil:
		return fmt.Errorf("%w:%s", workspace.ErrRepoDelete, err)
	}

	if _, err := r.Client.Collection(CollectionName).Doc(string(id)).Delete(ctx); err != nil {
		return fmt.Errorf("%w:%s", workspace.ErrRepoDelete, err)
	}

	return nil
}
