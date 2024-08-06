//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"cloud.google.com/go/logging"
	"github.com/markkurossi/go-libs/fn"
)

var (
	mux        *http.ServeMux
	projectID  string
	httpClient *http.Client
	blacklist  []Labels
	logInfo    *log.Logger
	logError   *log.Logger
)

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	mux.HandleFunc("/dns-query", DNSQuery)

	if !testing.Testing() {
		id, err := fn.GetProjectID()
		if err != nil {
			Fatalf("fn.GetProjectID: %s", err)
		}
		projectID = id

		// Create a logger client.
		ctx := context.Background()
		client, err := logging.NewClient(ctx, projectID)
		if err != nil {
			Fatalf("logging.NewClient: %v", err)
		}
		lg := client.Logger("NAP")
		logInfo = lg.StandardLogger(logging.Info)
		logError = lg.StandardLogger(logging.Error)
	}

	httpClient = new(http.Client)

	data, err := assets.ReadFile("data/default.bl")
	if err != nil {
		Fatalf("assets.ReadFile: %v", err)
	}
	blacklist, err = ParseBlacklist(data)
	if err != nil {
		Fatalf("ParseBlacklist: %v", err)
	}
}

// NAP implements the Google Cloud Functions entrypoint.
func NAP(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

// Errorf returns an HTTP error.
func Errorf(w http.ResponseWriter, code int, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	logError.Println(msg)
	http.Error(w, msg, code)
}

// Fatalf prints a fatal error and exits the program.
func Fatalf(format string, a ...interface{}) {
	log.Fatalf(format, a...)
}
