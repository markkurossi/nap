//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/markkurossi/go-libs/fn"
)

var (
	mux        *http.ServeMux
	projectID  string
	httpClient *http.Client
)

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	mux.HandleFunc("/dns-query", DNSQuery)

	if !testing.Testing() {
		id, err := fn.GetProjectID()
		if err != nil {
			Fatalf("fn.GetProjectID: %s\n", err)
		}
		projectID = id
	}

	httpClient = new(http.Client)
}

// NAP implements the Google Cloud Functions entrypoint.
func NAP(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

// Errorf returns an HTTP error.
func Errorf(w http.ResponseWriter, code int, format string, a ...interface{}) {
	http.Error(w, fmt.Sprintf(format, a...), code)
}

// Fatalf prints a fatal error and exits the program.
func Fatalf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
