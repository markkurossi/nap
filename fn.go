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
	"strings"
	"testing"

	"cloud.google.com/go/logging"
	"github.com/markkurossi/go-libs/fn"
)

var (
	mux        *http.ServeMux
	projectID  string
	httpClient *http.Client
	blacklists = make(map[string][]Labels)
	logInfo    *log.Logger
	logWarning *log.Logger
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
		logWarning = lg.StandardLogger(logging.Warning)
		logError = lg.StandardLogger(logging.Error)
	}

	httpClient = new(http.Client)

	entries, err := assets.ReadDir("data")
	if err != nil {
		Fatalf("assets.ReadDir: %v", err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".bl") {
			logWarning.Printf("unknown data file: %s", name)
			continue
		}

		data, err := assets.ReadFile("data/" + name)
		if err != nil {
			Fatalf("assets.ReadFile: %v", err)
		}
		blacklist, err := ParseBlacklist(data)
		if err != nil {
			Fatalf("ParseBlacklist: %v", err)
		}
		blacklists[name[:len(name)-3]] = blacklist
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
