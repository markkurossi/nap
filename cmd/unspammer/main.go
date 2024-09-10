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
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/markkurossi/nap/blacklist"
	"github.com/markkurossi/nap/handlers"
	"github.com/markkurossi/nap/pki"
)

var (
	bl           *blacklist.Blacklist
	certificates = make(map[string]*tls.Certificate)
)

func main() {
	blName := flag.String("blacklist", "", "DNS blacklist")
	caName := flag.String("ca", "", "The name of the CA")
	createCA := flag.Bool("create-ca", false, "Create CA")
	addr := flag.String("addr", ":443", "Address to listen")
	flag.Parse()

	log.SetFlags(0)

	if len(*blName) == 0 {
		log.Fatal("Blacklist name not specified")
	}
	err := readBlacklist(*blName)
	if err != nil {
		log.Fatal(err)
	}

	if len(*caName) == 0 {
		log.Fatal("CA name not specified")
	}

	var ca *pki.CA

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

func readBlacklist(list string) error {
	f, err := os.Open(list)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	bl, err = blacklist.ParseData(data)
	return err
}

var cors = map[string]string{
	"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Max-Age":           "1728000",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "*",
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s\n", r.Method, r.URL.Path)

	if false {
		for k, values := range r.Header {
			for _, v := range values {
				fmt.Printf(" - %v: %v\n", k, v)
			}
		}
	}

	for k, v := range cors {
		w.Header().Set(k, v)
	}

	if r.Method == "OPTIONS" {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		return
	}

	entry := bl.Match(r.Host)
	if entry == nil || entry.ProxyCmd == blacklist.ProxyBlock {
		handlers.Hello(w, r)
	} else if entry.ProxyCmd == blacklist.ProxyVAST {
		w.Header().Set("Content-Type", "text/xml")

		q := r.URL.Query()
		switch q.Get("resp") {
		case "vast4":
			log.Printf(" - returning vast4\n")
			fmt.Fprint(w, emptyVAST)

		case "vmap1+vast4":
			log.Printf(" - returning vmap1+vast4\n")
			fmt.Fprint(w, emptyVASTVMAP)

		default:
			log.Printf(" - unknown resp: %s\n", q.Get("resp"))
			fmt.Printf(" - q: %s\n", r.RequestURI)
			for k, v := range q {
				fmt.Printf(" - %s=%v\n", k, v)
			}
		}
	} else {
		proxy(w, r)
	}
}
