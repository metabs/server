package main

import (
	"flag"
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
		FeatureContext(s)
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

func FeatureContext(s *godog.Suite) {
	features.ServerIsUpAndRunning(s)
}
