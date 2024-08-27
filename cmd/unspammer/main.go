//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/markkurossi/nap/pki"
)

var (
	certificates = make(map[string]*tls.Certificate)
)

func main() {
	caName := flag.String("ca", "", "The name of the CA")
	createCA := flag.Bool("create-ca", false, "Create CA")
	addr := flag.String("addr", ":443", "Address to listen")
	flag.Parse()

	if len(*caName) == 0 {
		log.Fatal("CA name not specified")
	}

	var ca *pki.CA
	var err error

	if *createCA {
		ca, err = pki.CreateCA(*caName)
	} else {
		ca, err = pki.OpenCA(*caName)
	}
	if err != nil {
		log.Fatal(err)
	}
	eePriv, eePub, err := ca.CreateEEKey()
	if err != nil {
		log.Fatal(err)
	}
	eeTmpl := &x509.Certificate{}

	tlsConfig := &tls.Config{
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			fmt.Printf("ServerName: %v\n", info.ServerName)
			tlsCert, ok := certificates[info.ServerName]
			if !ok {
				eeTmpl.Subject = pkix.Name{
					CommonName: info.ServerName,
				}
				eeTmpl.DNSNames = []string{info.ServerName}

				cert, err := ca.CreateCertificate(eeTmpl, eePub)
				if err != nil {
					return nil, err
				}
				tlsCert = &tls.Certificate{
					Certificate: [][]byte{
						cert.Raw,
						ca.Cert.Raw,
					},
					PrivateKey: eePriv,
					Leaf:       cert,
				}
				certificates[info.ServerName] = tlsCert
			}
			return tlsCert, nil
		},
	}

	s := &http.Server{
		Addr:           *addr,
		Handler:        http.HandlerFunc(handler),
		TLSConfig:      tlsConfig,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServeTLS("", ""))
}

var cors = map[string]string{
	"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Max-Age":           "1728000",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "*",
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s\n", r.Method, r.URL.Path)
	if false {
		for k, values := range r.Header {
			for _, v := range values {
				fmt.Printf(" - %v: %v\n", k, v)
			}
		}
	}

	if r.Method == "OPTIONS" {

		if false {
			setHdr(r, w, "Access-Control-Request-Private-Network",
				"Private-Network")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods",
				"GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "240")
			w.WriteHeader(204)
		}
		for k, v := range cors {
			w.Header().Set(k, v)
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/ad/") {
		for k, v := range cors {
			w.Header().Set(k, v)
		}
		w.Header().Set("Content-Type", "text/xml")

		q := r.URL.Query()
		switch q.Get("resp") {
		case "vast4":
			fmt.Printf(" - returning vast4\n")
			fmt.Fprint(w, emptyVAST)

		case "vmap1+vast4":
			fmt.Printf(" - returning vmap1+vast4\n")
			fmt.Fprint(w, emptyVASTVMAP)

		default:
			fmt.Printf(" - unknown resp: %s\n", q.Get("resp"))
			fmt.Printf(" - q: %s\n", r.RequestURI)
			for k, v := range q {
				fmt.Printf(" - %s=%v\n", k, v)
			}
		}
	} else {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.String()))
	}
}

func setHdr(r *http.Request, w http.ResponseWriter, req, resp string) {
	val := r.Header.Get(req)
	if len(val) > 0 {
		w.Header().Set("Access-Control-Allow-"+resp, val)
	}
}
