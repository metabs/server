package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"flag"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/db"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/jwt"
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
		dbc, err := db.New(context.Background())
		if err != nil {
			panic(err)
		}
		sv, err := jwt.New()
		if err != nil {
			panic(err)
		}
		FeatureContext(s, sv, dbc)
	}, godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
		//Randomize:     time.Now().UTC().UnixNano(),
		StopOnFailure: stopOnFailure,
	})

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func FeatureContext(s *godog.Suite, sv *jwt.SignerVerifier, db *firestore.Client) {
	if _, err := db.Collection("Ping").Doc("0").Set(context.Background(), map[string]bool{"1": true}); err != nil {
		panic(err)
	}
	s.AfterScenario(func(i interface{}, err error) {
		if err := deleteCollections(db, "workspace"); err != nil {
			panic(err)
		}
		if err := deleteCollections(db, "customer"); err != nil {
			panic(err)
		}
	})
	features.ServerIsUpAndRunning(s)
	features.WorkspaceAPIs(s, sv, db)
}

// copied from google doc :)
func deleteCollections(client *firestore.Client, collectionName string ) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	docs := client.Collection(collectionName).Documents(ctx)

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
