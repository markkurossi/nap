//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var googlebotHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
}

func proxy(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)
	log.Printf(" - proxy => %s\n", url)

	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		log.Print(err)
		return
	}

	// Google only ever indexes one page worth of content for a
	// specific URL. Even if you could get Googlebot to change the
	// language, every time it recrawled the URL with different
	// cookies, it would overwrite its indexing data for other
	// languages.

	// The only way to have Google index content in multiple languages
	// is to have different URLs for each language. See How should I
	// structure my URLs for both SEO and localization?

	// In any case, Googlebot doesn't use cookies that persist between
	// pages. Googlebot fetches each page with an empty cookie jar. It
	// does so because users that click through to your site from
	// Google are not going to start with any cookies set for your
	// site.

	for k, values := range r.Header {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	req.Header.Del("Cookie")
	if true {
		for k, v := range googlebotHeaders {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	for k, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Print(err)
		return
	}
}
