package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"flag"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/db"
	"google.golang.org/api/iterator"
	"os"
	"testing"
	"time"

	"github.com/unprogettosenzanomecheforseinizieremo/server/tests/features"

	"github.com/cucumber/godog"
)

func TestMain(m *testing.M) {
	var runCucumberTests bool
	var stopOnFailure bool
	flag.BoolVar(&runCucumberTests, "cucumber", false, "Set this flag if you want to run godog BDD tests")
	flag.BoolVar(&stopOnFailure, "stop-on-failure", false, "Stop processing on first failed scenario.. Flag is passed to godog")
	flag.Parse()

	if !runCucumberTests {
		os.Exit(0)
	}

	status := godog.RunWithOptions("App", func(s *godog.Suite) {
		db, err := db.New(context.Background())
		if err != nil {
			panic(err)
		}
		FeatureContext(s, db)
	}, godog.Options{
		Format:        "pretty",
		Paths:         []string{"features"},
		Randomize:     time.Now().UTC().UnixNano(),
		StopOnFailure: stopOnFailure,
	})

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func FeatureContext(s *godog.Suite, db *firestore.Client) {
	if _, err := db.Collection("Ping").Doc("0").Set(context.Background(), map[string]bool{"1": true}); err != nil {
		panic(err)
	}
	s.AfterScenario(func(i interface{}, err error) {
		if err := deleteCollection(db); err != nil {
			panic(err)
		}
	})
	features.ServerIsUpAndRunning(s)
	features.WorkspaceAPIs(s, db)
}

// copied from google doc :)
func deleteCollection(client *firestore.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	docs := client.Collection("workspace").Documents(ctx)

	numDeleted := 0
	batch := client.Batch()
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		batch.Delete(doc.Ref)
		numDeleted++
	}

	// If there are no documents to delete,
	// the process is over.
	if numDeleted == 0 {
		return nil
	}

	_, err := batch.Commit(ctx)
	return err

}
