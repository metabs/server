package db

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
)

var projectID = os.Getenv("FIRESTORE_PROJECT_ID")

// New Create a database connection using the environment variable to define the database driver and url
// it returns an error when an error occurs establishing the connection
func New(ctx context.Context) (*firestore.Client, error) {

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return client, nil
}
