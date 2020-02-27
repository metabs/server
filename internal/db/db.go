package db

import (
	"context"
	"google.golang.org/api/option"
	"os"

	"cloud.google.com/go/firestore"
)

var projectID = os.Getenv("FIRESTORE_PROJECT_ID")
var saPath = os.Getenv("SA_PATH")

// New Create a database connection using the environment variable to define the database driver and url
// it returns an error when an error occurs establishing the connection
func New(ctx context.Context) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(saPath)
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		return nil, err
	}

	return client, nil
}
