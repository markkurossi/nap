//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package handlers

import (
	"net/http"
)

// Hello implements handler for healthcheck calls.
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' width='48' height='48' viewBox='0 0 16 16'><text x='0' y='14'>🚫</text></svg>" >
    <title>NAP</title>
    <style>
        body {
            font-family: monospace;
        }
    </style>
  </head>
  <body>
    <h1>&#x1f6ab; No Advertising Please &#x1f6ab;</h1>
  </body>
</html>
`))
}
